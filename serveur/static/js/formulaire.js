function updatePourcentage(val, idValue) {
    document.getElementById(idValue).innerText = val + '%';
}

function afficher(idElemChecked, idAAfficher, idACacher) {
    var elemChecked = document.getElementById(idElemChecked);
    var aAfficher = document.getElementById(idAAfficher);
    var aCacher = document.getElementById(idACacher);

    if (elemChecked.checked) {
        aAfficher.style.display = 'block';
        aCacher.style.display = 'none'
    } else {
        aAfficher.style.display = 'none';
    }
}

function sendData(event) {
    event.preventDefault();

    const id_simulation = document.getElementById('id_simulation').value;
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
            strat_avant = -1;
        }

        if(document.getElementById('placesReservees').checked) { //ON A CHOISI PLACES RESERVEES
            pourcentage_places_avant = document.getElementById('pourcentage_places_sans').value;
        } else {
            pourcentage_places_avant = 0;
        }

        // On met tout à vide pour la partie "Après"
        objectif = 0
        type_recrutement_apres = -1
        strat_apres = -1
        pourcentage_places_apres = 0

    } else {

        /*AVEC OBJECTIF : ON MET DANS 'AVANT'*/

        objectif = document.getElementById('objectif').value;

        /*RECRUTEMENT AVANT*/
        type_recrutement_avant = document.querySelector('input[name="type_recrutement_avant"]:checked').value;
        if(document.getElementById('competencesAvant').checked) { //ON A CHOISI 'COMPETENCES EGALES
            strat_avant = document.querySelector('input[name="strat_avant"]:checked').value;
        } else {
            strat_avant = -1;
        }

        if(document.getElementById('placesReserveesAvant').checked) { //ON A CHOISI PLACES RESERVEES
            pourcentage_places_avant = document.getElementById('pourcentage_places_avant').value;
        } else {
            pourcentage_places_avant = 0;
        }

        /*RECRUTEMENT APRES*/
        type_recrutement_apres = document.querySelector('input[name="type_recrutement_apres"]:checked').value;
        if(document.getElementById('competencesApres').checked) { //ON A CHOISI 'COMPETENCES EGALES
            strat_apres = document.querySelector('input[name="strat_apres"]:checked').value;
        } else {
            strat_apres = -1;
            console.log("ici")
        }

        if(document.getElementById('placesReserveesApres').checked) { //ON A CHOISI PLACES RESERVEES
            pourcentage_places_apres = document.getElementById('pourcentage_places_apres').value;
        } else {
            pourcentage_places_apres = 0;
        }
    }

    const formData = {
      id_simulation : id_simulation,
      nb_employes : Number(nb_employes),
      nb_annees : Number(nb_annees),
      pourcentage_femmes : Number(pourcentage_femmes)/100,
      objectif : Number(objectif)/100,
      strat_avant : Number(strat_avant),
      strat_apres : Number(strat_apres),
      type_recrutement_avant : Number(type_recrutement_avant),
      type_recrutement_apres : Number(type_recrutement_apres),
      pourcentage_places_avant : Number(pourcentage_places_avant)/100,
      pourcentage_places_apres : Number(pourcentage_places_apres)/100
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
    .then((response) => {
      if (!response.ok) {
      throw new Error(`HTTP error! Status: ${response.status}`);
      } else {
          response.json().then(data => console.log(data))
      }
      })
      
  }
