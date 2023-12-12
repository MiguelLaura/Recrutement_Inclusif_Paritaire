let id_simulation = 0

function updatePourcentage(val, idValue) {
    document.getElementById(idValue).innerText = val + '%';
}

function afficher(idElemChecked, idAAfficher, idACacher) {
    var elemChecked = document.getElementById(idElemChecked);
    var aAfficher = document.getElementById(idAAfficher);
    var aCacher = document.getElementById(idACacher);

    var inputsAAfficher = aAfficher.querySelectorAll('input');
    var inputsACacheer = aCacher.querySelectorAll('input');

    if (elemChecked.checked) {
        aAfficher.style.display = 'block';
        inputsAAfficher.forEach(function(inputsAAfficher) {
            inputsAAfficher.required = true;
        });
        aCacher.style.display = 'none'
        inputsACacheer.forEach(function(inputsACacheer) {
            inputsACacheer.required = false;
        });
    }
}

function sendData(event) {
    event.preventDefault();

    const nb_employes = document.getElementById('nb_employes').value;
    const nb_annees = document.getElementById('nb_annees').value;
    const pourcentage_femmes = document.getElementById('pourcentage_femmes').value;

    var strat_avant;
    var type_recrutement_avant;
    var pourcentage_places_avant;
    var objectif;
    var strat_apres;
    var type_recrutement_apres;
    var pourcentage_places_apres;
    var objectif;

    /*SANS OBJECTIF : ON MET DANS 'AVANT'*/

    if(document.getElementById('nonObjectif').checked) {
        type_recrutement_avant = document.querySelector('input[name="type_recrutement_sans"]:checked').value;
        if(document.getElementById('competences').checked) { //ON A CHOISI 'COMPETENCES EGALES
            strat_avant = document.querySelector('input[name="strat_sans"]:checked').value;
        } else {
            strat_avant = 0; //StratVide dans l'énumération
        }

        if(document.getElementById('placesReservees').checked) { //ON A CHOISI PLACES RESERVEES
            pourcentage_places_avant = document.getElementById('pourcentage_places_sans').value;
            pourcentage_places_avant = Number(pourcentage_places_avant)/100
        } else {
            pourcentage_places_avant = -1;
        }

        // On met tout à vide pour la partie "Après"
        objectif = -1
        type_recrutement_apres = 0
        strat_apres = 0
        pourcentage_places_apres = -1

    } else {

        /*AVEC OBJECTIF : ON MET DANS 'AVANT'*/

        objectif = document.getElementById('objectif').value;
        objectif = Number(objectif)/100,

        /*RECRUTEMENT AVANT*/
        type_recrutement_avant = document.querySelector('input[name="type_recrutement_avant"]:checked').value;
        if(document.getElementById('competencesAvant').checked) { //ON A CHOISI 'COMPETENCES EGALES
            strat_avant = document.querySelector('input[name="strat_avant"]:checked').value;
        } else {
            strat_avant = 0;
        }

        if(document.getElementById('placesReserveesAvant').checked) { //ON A CHOISI PLACES RESERVEES
            pourcentage_places_avant = document.getElementById('pourcentage_places_avant').value;
            pourcentage_places_avant = Number(pourcentage_places_avant)/100
        } else {
            pourcentage_places_avant = -1;
        }

        /*RECRUTEMENT APRES*/
        type_recrutement_apres = document.querySelector('input[name="type_recrutement_apres"]:checked').value;
        if(document.getElementById('competencesApres').checked) { //ON A CHOISI 'COMPETENCES EGALES
            strat_apres = document.querySelector('input[name="strat_apres"]:checked').value;
        } else {
            strat_apres = 0;
        }

        if(document.getElementById('placesReserveesApres').checked) { //ON A CHOISI PLACES RESERVEES
            pourcentage_places_apres = document.getElementById('pourcentage_places_apres').value;
            pourcentage_places_apres = Number(pourcentage_places_apres)/100
        } else {
            pourcentage_places_apres = -1;
        }
    }

    id_simulation +=1

    const formData = {
      id_simulation : "id_simulation_"+id_simulation,
      nb_employes : Number(nb_employes),
      nb_annees : Number(nb_annees),
      pourcentage_femmes : Number(pourcentage_femmes)/100,
      objectif : objectif,
      strat_avant : Number(strat_avant),
      strat_apres : Number(strat_apres),
      type_recrutement_avant : Number(type_recrutement_avant),
      type_recrutement_apres : Number(type_recrutement_apres),
      pourcentage_places_avant : pourcentage_places_avant,
      pourcentage_places_apres : pourcentage_places_apres
    };

    console.log(formData)

    // Envoi des données au serveur en utilisant la méthode fetch avec la méthode POST
    fetch('http://localhost:8080/new_simulation', {
      method: 'POST',
      headers: {
          "Content-Type": "application/json",
          "Accept": "application/json",
          "Access-Control-Allow-Origin":'*'
      },
      body: JSON.stringify(formData)
    })
    .then(response => {
        if (!response.ok) {
          return response.text();
        } else {
          return response.json();
        }
      })
      .then(data => {
        console.log(data);
        var respValue = document.getElementById("responseFormValue");
        if(typeof data === 'object') { 
            respValue.style.color = "black";
            respValue.innerText = `Simulation créée ! [ID : ${data.id_simulation}]`;
        } else { //Erreur
            respValue.style.color = "red";
            respValue.innerText = data;
        }
      })
      .catch(error => {
        console.error('Erreur HTTP:', error);
      });
      
  }
