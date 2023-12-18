// -------------------------
// Chart

const dataBenef = {
    label: "Bénéfices",
    data: [2500, 1900, 10000, 4950, 4950, 3970, 5300],
    fill: false,
    borderColor: 'rgb(130, 174, 210)',
}

const dataParite = {
    label: "Parité",
    data: [0, 600, 600, 1000, 500, 500, 1300],
    fill: false,
    borderColor: 'rgb(231, 54, 56)',
}

const ctx = document.getElementById('sim-graph');

const graph = new Chart(ctx, {
  type: 'line',
  data: {
    labels: [0, 5, 10, 15, 20, 25, 30],
    datasets: [dataBenef, dataParite],
  }
});

// Fonction de démo pour montrer comment ajouter une nouvelle valeur au graphe
function ajoutDonneeGraph(donne1,donne2) {
    // Ajout d'un x
    graph.data.labels.push(graph.data.labels[graph.data.labels.length-1] + 5);
    // Ajout des y
    graph.data.datasets[0].data.push(donne1);
    graph.data.datasets[1].data.push(donne2);
    // Mise à jour du graphe
    graph.update();
}