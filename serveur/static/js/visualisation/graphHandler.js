// -------------------------
// Chart

let titles = {
    "Bénéfices" : "Bénéfices produit par l'entreprise par année",
    "Parité" : "Pourcentage de femmes dans l'entreprise",
    "Compétences" : "Moyenne des compétences des employé-es (sur 10)",
    "Santé mentale" : 'Moyenne de la santé mentale des employé-es (sur 100)'
}

class Graph {
    constructor(parent) {
        this.xIncr = 1;
        this.xs = [];
        this.theGraphs = [];
        this.theLimits = {};

        this.renderedLimits = {}
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
                    }
                }
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
     * @param {Array} color - couleur du nouveau graph au format rgb : [R, G, B]
     */
    addNewGraph(label, color = undefined) {
        if (color === undefined) {
            color = randomRGBColor()
        }
        const newY = {
            label: label,
            data: [],
            fill: false,
            borderColor: `rgb(${color[0]}, ${color[1]}, ${color[2]})`,
            title: {
                display: true,
                text: titles[label],
                color: '#191',
            }
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
     *      labelBkgColor - La couleur du fond du label (defaut: "red")
     *      width - La largeur de la ligne (en px) (defaut: 3)
     *      showLabel - Est-ce que le label est visible ou non (defaut : true)
     */
    addHorizontalLine(graphLabel, lineLabel, value, { lineColor = "black", labelBkgColor = "red", width = 3, showLabel = true } = {}) {
        if (Object.keys(this.theLimits).indexOf(graphLabel) === -1) {
            throw new Error("Ne peut pas ajouter de limites à un graphe qui n'existe pas.");
        }

        const limit = {
            lineLabel: lineLabel,
            type: 'line',
            borderColor: lineColor,
            borderWidth: width,
            label: {
                backgroundColor: labelBkgColor,
                content: lineLabel,
                display: showLabel
            },
            scaleID: 'y',
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
        this.graph.data.datasets = [];
        emptyObject(this.renderedLimits);
        for (const graph of this.theGraphs) {
            if (graphNames.indexOf(graph.label) !== -1) {
                this.graph.data.datasets.push(graph);
                for (const limit of this.theLimits[graph.label]) {
                    this.renderedLimits[`${graph.label}_${limit.lineLabel}`] = limit;
                }
            }
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
}

function randomRGBColor() {
    return [
        Math.floor(Math.random() * 256),
        Math.floor(Math.random() * 256),
        Math.floor(Math.random() * 256)
    ]
}

function emptyObject(obj) {
    for (let keyName in obj) {
        delete obj[keyName]
    }
    return obj
}