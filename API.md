
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
| `objectif` | `float` | **Required**. Objectif pourcentage min de femmes dans l'entreprise. Peut-être -1 (pas d'objectif) |
| `strat_avant` | `int` | **Required**. Vide = 0 / PrioHommes = 1 / PrioFemmes = 2 / Hasard = 3 / |
| `strat_apres` | `int` | **Required**. Vide = 0 / PrioHommes = 1 / PrioFemmes = 2 / Hasard = 3 / |
| `type_recrutement_avant` | `int` | **Required**.  Vide = 0 / CompetencesEgales = 1 / PlacesReservees = 2 / |
| `type_recrutement_apres` | `int` | **Required**. Vide = 0 / CompetencesEgales = 1 / PlacesReservees = 2 / |
| `pourcentage_places_avant` | `float` | **Required**. Pourcentages de places réservées dans le recrutement. Peut-être -1 (pas de places réservées) |
| `pourcentage_places_apres` | `float` | **Required**. Pourcentages de places réservées dans le recrutement. Peut-être -1 (pas de places réservées) |

- Code retour

| Code retour | Signification |
|-------------|---------------|
| `201`       | simulation créée     |
| `400`       | mauvaise requête   |

- Objet `JSON` renvoyé (si `201`)

| propriété  | type | exemple de valeurs possibles                                  |
|------------|-------------|-----------------------------------------------------|
| `simulation-id`    | `string` | `"id_simulation_1"` |





