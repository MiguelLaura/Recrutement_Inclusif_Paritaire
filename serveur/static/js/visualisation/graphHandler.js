// -------------------------
// Chart

class Graph {
    constructor(parent) {
        this.xs = [];
        this.theGraphs = [];
        this.graph = new Chart(parent, {
            type: 'line',
            data: {
                labels: this.xs,
                datasets: []
            }
        });
        this.xIncr = 1;
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
    }

    /**
     * Selectionne les graphes qui ont le label indiqué dans la liste d'arguments
     * @param  {...string} graphNames Les labels des graphs que l'on souhaite afficher
     */
    selectGraphs(...graphNames) {
        this.graph.data.datasets = [];
        for(const graph of this.theGraphs) {
            if(graphNames.indexOf(graph.label) !== -1) {
                this.graph.data.datasets.push(graph);
            }
        }
    }

    /**
     * Ajoute aux graphs une nouvelle valeur en Y
     * @param  {...number} allGraphData Les valeurs à attribuer aux graphs, dans le même ordre que lorsque addNewGraph a été appelé
     */
    addData(...allGraphData) {
        if(this.xs.length === 0) {
            this.xs.push(0);
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