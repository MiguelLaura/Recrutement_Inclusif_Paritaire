
## API Reference

### Créer une nouvelle simulation

```http
  POST /localhost:8080/new_simulation
```

- Requête : `POST`
- Objet `JSON` envoyé


| Paramètres | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `id_simulation` | `string` | **Required**. ID de l'entreprise créée |
| `nb_employes` | `int` | **Required**. Nombre d'employés total dans l'entreprise |
| `nb_annees` | `int` | **Required**. Nombre d'années pour la simulation |
| `pourcentage_femmes` | `float` | **Required**. Pourcentage de femmes au début |
| `objectif` | `float` | **Required**. Objectif pourcentage min de femmes dans l'entreprise. Peut-être 0% (pas d'objectif) |
| `strat_avant` | `int` | **Required**. PrioHommes = 0 / PrioFemmes = 1 / Hasard = 2 / Vide = -1 |
| `strat_apres` | `int` | **Required**. PrioHommes = 0 / PrioFemmes = 1 / Hasard = 2 / Vide = -1 |
| `type_recrutement_avant` | `int` | **Required**. CompetencesEgales = 0 / PlacesReservees = 1 / Vide = -1 |
| `type_recrutement_apres` | `int` | **Required**. CompetencesEgales = 0 / PlacesReservees = 1 / Vide = -1 |
| `pourcentage_places_avant` | `float` | **Required**. Pourcentages de places réservées dans le recrutement |
| `pourcentage_places_apres` | `float` | **Required**. Pourcentages de places réservées dans le recrutement |

- Code retour

| Code retour | Signification |
|-------------|---------------|
| `201`       | simulation créée     |
| `400`       | mauvaise requête   |

- Objet `JSON` renvoyé (si `201`)

| propriété  | type | exemple de valeurs possibles                                  |
|------------|-------------|-----------------------------------------------------|
| `simulation-id`    | `string` | `"simulation-1"` |





