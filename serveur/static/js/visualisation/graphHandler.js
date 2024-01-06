// -------------------------
// Chart

class Graph {
    constructor(parent) {
        this.xIncr = 1;
        this.parent = parent;
        this.xs = [];
        this.theGraphs = [];
        this.theLimits = {};

        // Configs
        this.currTitle = {
            display: false,
            text: "",
            color: "#191"
        };
        this.renderedLimits = {}
        this.scales = {}

        // Graph
        this.graph = new Chart(parent, {
            type: 'line',
            data: {
                labels: this.xs,
                datasets: []
            },
            options: {
                plugins: {
                    annotation: {
                        annotations: this.renderedLimits
                    },
                    title: this.currTitle
                },
                scales: this.scales
            }
        });
    }

    setIncrement(newIncr) {
        if (newIncr < 1) {
            throw new Error("La quantité d'incrémentation du x du graphe ne peut pas être < 1")
        }
        this.xIncr = newIncr;
    }

    /**
     * Ajoute un nouveau graph
     * @param {string} label - Nom du nouveau graph 
     * @param {object}  - Paramètres optionnels :
     *      color - couleur du nouveau graph au format rgb : [R, G, B] (défaut: aléatoire)
     *      title - titre du graph
     *      beginAtZero - Est-ce que le graph commence à 0 en Y (défaut : true)
     */
    addNewGraph(label, { color = undefined, title = "", beginAtZero = true, min = undefined, max = undefined } = {}) {
        if (color === undefined) {
            color = randomRGBColor()
        }

        const newY = {
            label: label,
            data: [],
            fill: false,
            borderColor: `rgb(${color[0]}, ${color[1]}, ${color[2]})`,
            title: title,
            yAxisID: this._addYAxes({ beginAtZero: beginAtZero, min: min, max: max, color: color })
        }

        this.theGraphs.push(newY);
        this.theLimits[label] = [];
    }

    /**
     * Ajoute une ligne horizontale sur le graphe
     * @param {string} graphLabel - Le nom du graph auquel la ligne est rattachée 
     *      (si le graphe n'existe pas, une exception est générée)
     * @param {string} lineLabel - Le label de la ligne
     * @param {float} value - La valeur en y de la ligne
     * @param {object}  - Paramètres optionnels :
     *      lineColor - La couleur de la ligne (defaut: "black")
     *      labelBkgColor - La couleur du fond du label (defaut: undefined == couleur du graphe)
     *      width - La largeur de la ligne (en px) (defaut: 3)
     *      showLabel - Est-ce que le label est visible ou non (defaut : true)
     */
    addHorizontalLine(graphLabel, lineLabel, value, { lineColor = "black", labelBkgColor = undefined, width = 3, showLabel = true } = {}) {
        if (Object.keys(this.theLimits).indexOf(graphLabel) === -1) {
            throw new Error("Ne peut pas ajouter de limites à un graphe qui n'existe pas.");
        }

        const theGraph = this._getGraph(graphLabel);

        const limit = {
            lineLabel: lineLabel,
            type: 'line',
            borderColor: lineColor,
            borderWidth: width,
            label: {
                backgroundColor: labelBkgColor === undefined ? theGraph.borderColor : labelBkgColor,
                content: lineLabel,
                display: showLabel
            },
            scaleID: theGraph.yAxisID,
            value: value
        };

        this.theLimits[graphLabel].push(limit);
        if (this.isGraphActive(graphLabel)) {
            this.renderedLimits[`${graphLabel}_${lineLabel}`] = limit;
        }
    }

    /**
     * Selectionne les graphes qui ont le label indiqué dans la liste d'arguments
     * @param  {...string} graphNames Les labels des graphs que l'on souhaite afficher
     */
    selectGraphs(...graphNames) {
        // Resets graph
        this._hideAxes();
        this.currTitle.text = "";
        this.currTitle.display = false;
        emptyObject(this.renderedLimits);

        // Select the axis
        let graphCnt = 0;
        for (const graph of this.theGraphs) {
            if (graphNames.indexOf(graph.label) !== -1) {
                const axes = this._getYAxes(graph.label);
                axes.display = true;
                axes.position = graphCnt === 0 ? "left" : "right";
                axes.grid.drawOnChartArea = graphCnt === 0 ? true : false;
                axes.ticks.color = graphNames.length === 1 ? "black" : graph.borderColor;
                graphCnt++;
            }
        }

        this._changeGraph();
        
        // Shows new graphs
        for (const graph of this.theGraphs) {
            if (graphNames.indexOf(graph.label) !== -1) {
                this.graph.data.datasets.push(graph);

                if(graph.title.trim() != "" && graphNames.length == 1) {
                    this.currTitle.text = graph.title.trim();
                    this.currTitle.display = true;
                }

                for (const limit of this.theLimits[graph.label]) {
                    this.renderedLimits[`${graph.label}_${limit.lineLabel}`] = limit;
                }
            }
        }

        if(graphNames.length > 1) {
            this.currTitle.text = `graphe multiple de ${graphNames.join(", ")}`;
            this.currTitle.display = true;
        }
    }

    /**
     * Ajoute aux graphs une nouvelle valeur en Y
     * @param  {...number} allGraphData Les valeurs à attribuer aux graphs, dans le même ordre que lorsque addNewGraph a été appelé
     */
    addData(...allGraphData) {
        if (this.xs.length === 0) {
            this.xs.push(1);
        } else {
            this.xs.push(this.xs[this.xs.length - 1] + this.xIncr);
        }

        if (allGraphData.length !== this.theGraphs.length) {
            throw new Error("Pas autant de données de de graphs")
        }

        for (const graphIdx in this.theGraphs) {
            this.theGraphs[graphIdx].data.push(allGraphData[graphIdx])
        }

        this.render();
    }

    isGraphActive(graphLabel) {
        return this.graph.data.datasets.some(graph => graph.label === graphLabel);
    }

    render() {
        this.graph.update();
    }

    reset() {
        this.xs.length = 0;
        for (const graphIdx in this.theGraphs) {
            this.theGraphs[graphIdx].data.length = 0;
        }
    }

    _getGraph(graphLabel) {
        for(const graph of this.theGraphs) {
            if(graph.label === graphLabel) {
                return graph;
            }
        }
        return undefined;
    }

    _getYAxes(graphLabel) {
        const axesLabel = this._getGraph(graphLabel).yAxisID;

        return this.scales[axesLabel];
    }

    _addYAxes({ beginAtZero = true, position = 'left', display = false, min = undefined, max = undefined, color = undefined } = {}) {
        const newAxes = {
            type: 'linear',
            display: display,
            beginAtZero: beginAtZero,
            position: position,
            grid: {
                drawOnChartArea: false,
            }
        };

        if(min !== undefined) {
            newAxes["min"] = min;
        }

        if(max !== undefined) {
            newAxes["max"] = max;
        }

        if(color !== undefined) {
            newAxes["ticks"] = {
                color: color
            }
        }

        const axesId = `y${Object.keys(this.scales).length}`;
        this.scales[axesId] = newAxes;

        return axesId;
    }

    _hideAxes() {
        for(const axes in this.scales) {
            this.scales[axes].display = false;
        }
    }

    _changeGraph() {
        this.graph.destroy()
        this.graph = new Chart(this.parent, {
            type: 'line',
            data: {
                labels: this.xs,
                datasets: []
            },
            options: {
                plugins: {
                    annotation: {
                        annotations: this.renderedLimits
                    },
                    title: this.currTitle
                },
                scales: this.scales
            }
        });
    }
}

function randomRGBColor() {
    return [
        Math.floor(Math.random() * 256),
        Math.floor(Math.random() * 256),
        Math.floor(Math.random() * 256)
    ];
}

function emptyObject(obj) {
    for (let keyName in obj) {
        delete obj[keyName];
    }
    return obj;
}