// -------------------------
// Chart

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
        if(newIncr < 1) {
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
        if(color === undefined) {
            color = randomRGBColor()
        }
        const newY = {
            label: label,
            data: [],
            fill: false,
            borderColor: `rgb(${color[0]}, ${color[1]}, ${color[2]})`,
        }

        this.theGraphs.push(newY);
        this.theLimits[label] = [];
    }

    /**
     * Ajoute une ligne horizontale sur le graphe
     * @param {string} graphLabel - Le nom du graph auquel la ligne est rattachée
     * @param {string} limitLabel - Le label de la ligne
     * @param {float} limitValue - La valeur en y de la ligne
     * @param {*} limitLineColor - La couleur de la ligne
     * @param {*} limitBkgColor - La couleur du fond du label
     * @param {int} width - La largeur de la ligne
     */
    addHorizontalLine(graphLabel, limitLabel, limitValue, limitLineColor="black", limitBkgColor="red", width=3) {
        if(Object.keys(this.theLimits).indexOf(graphLabel) === -1) {
            throw new Error("Ne peut pas ajouter de limites à un graphe qui n'existe pas.");
        }

        const limit = {
            limitName: limitLabel,
            type: 'line',
            borderColor: limitLineColor,
            borderWidth: width,
            label: {
                backgroundColor: limitBkgColor,
                content: limitLabel,
                display: true
            },
            scaleID: 'y',
            value: limitValue
        };

        this.theLimits[graphLabel].push(limit);
        if(this.isGraphActive(graphLabel)) {
            this.renderedLimits[`${graphLabel}_${limitLabel}`] = limit;
        }
    }

    /**
     * Selectionne les graphes qui ont le label indiqué dans la liste d'arguments
     * @param  {...string} graphNames Les labels des graphs que l'on souhaite afficher
     */
    selectGraphs(...graphNames) {
        this.graph.data.datasets = [];
        emptyObject(this.renderedLimits);
        for(const graph of this.theGraphs) {
            if(graphNames.indexOf(graph.label) !== -1) {
                this.graph.data.datasets.push(graph);
                for(const limit of this.theLimits[graph.label]) {
                    this.renderedLimits[`${graph.label}_${limit.limitName}`] = limit;
                }
            }
        }
    }

    /**
     * Ajoute aux graphs une nouvelle valeur en Y
     * @param  {...number} allGraphData Les valeurs à attribuer aux graphs, dans le même ordre que lorsque addNewGraph a été appelé
     */
    addData(...allGraphData) {
        if(this.xs.length === 0) {
            this.xs.push(1);
        } else {
            this.xs.push(this.xs[this.xs.length-1] + this.xIncr);
        }

        if(allGraphData.length !== this.theGraphs.length) {
            throw new Error("Pas autant de données de de graphs")
        }

        for(const graphIdx in this.theGraphs) {
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
        for(const graphIdx in this.theGraphs) {
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
    for(let keyName in obj) {
        delete obj[keyName]
    }
    return obj
}