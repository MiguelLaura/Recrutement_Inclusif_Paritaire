**ATTENTION REMPLACER LES TEST PAR NUMERO + FAIRE LES A FAIRE**

# La parité en entreprise

## Contexte de réalisation du projet
Ce projet a été réalisé dans le cadre de l'UV IA04 enseignée à l'université de Technologie de Compiègne. Il a été réalisé par :
* Mathilde Lange
* Solenn Lenoir
* Nathan Menny
* Laura Miguel

## Sujet
Ce projet est une simulation d'une entreprise. On souhaite étudier la parité dans une entreprise. Nous nous sommes en particulier concentré.e.s sur le recrutement : c'est une étape systématique dans une entreprise qui peut influer sur la parité. Le but est donc de laisser l'utilisateur.ice choisir une façon d'effectuer son recrutement, de lancer la simulation et voir l'influence de ce recrutement sur le bénéfice de l'entreprise, le nombre de départs, etc.

## Problématique
Notre principale question est donc : quelles sont les conséquences d’un recrutement tenant compte de la parité sur la santé d’une entreprise ?

## Récupérer le projet

```bash
go get github.com/gorilla/websocket
```
**A FAIRE**

## Les entrées de l'utilisateur.ice
L'utilisateur.ice va pouvoir définir :
* le nombre d'employé.e.s de l'entreprise,
* la durée de la simulation (un pas de la simulation correspond à une année dans l'entreprise),
* le pourcentage initial de femmes dans l'entreprise,
* s'iel a un objectif de parité (défini sur l'interface comme un pourcentage de femmes à atteindre),
* le type de recrutement (s'iel a un objectif, il faut choisir un type de recrutement pour quand on est en dessous du pourcentage de femmes voulues, et un type de recrutement pour quand on est au-dessus du pourcentage de femmes voulues).

### Pourquoi un pourcentage de femmes à atteindre ?
Les entreprises peuvent vouloir atteindre une certaine parité (pour respecter une loi, favoriser l’innovation, etc.) et mettre en place des stratégies temporaires, notamment au niveau du recrutement. Ainsi, les manières de recruter ne seront pas les mêmes en dessous ou au-dessus du seuil défini.

### Places réservées ?
Pour cette stratégie, sur le nombre de personnes à recruter, on choisira de recruter un pourcentage fixe de femmes ou d'hommes (on prendra toujours les plus compétent.e.s dans cette population) puis, pour le reste des candidat.e.s, on recrutera en fonction des compétences seulement. Cette stratégie n'existe pas dans la réalité puisqu'il s'agit d'une discrimination de genre. En effet, on ne peut discriminer à l'embauche sur le genre que pour des cas particuliers, comme pour le cinéma ou mannequinat[<sup>test</sup>](https://analyseur.acompetenceegale.com/comment-eviter-discriminations-a-lembauche-selon-sexe/).

### Compétences égales ?
Pour cette stratégie, on recrute d’abord la personne la plus compétente. Si jamais deux personnes ont des compétences équivalentes, on choisira qui recruter en fonction de ce qui a été demandé par l’utilisateur.ice : iel choisit s'iel donne sa préférence à une femme, à un homme, ou s'iel n'a pas de préférence et prend un.e des candidat.e.s au hasard. C'est un type de recrutement qu'on peut appliquer à la vie réelle, mais uniquement en cas de candidatures comparables, en faveur du genre sous-représenté et en cas de dernier critère de départage[<sup>test</sup>](https://egaliteautravail.com/domaine/recrutement/).

## L'interface de simulation

### Le formulaire
Nous avons en premier temps un formulaire dans lequel l'utilisateur.ice entre les informations mentionnées dans la partie [Les entrées de l'utilisateur.ice](#les-entrées-de-lutilisateurice). Il y a également des informations sur la simulation qui correspondent aux parties précédentes.

### La simulation
**A FAIRE**

## Non pris en compte dans notre modélisation[<sup>test</sup>](https://infonet.fr/actualite/focus/parite-femme-homme-en-entreprise-7-pratiques-a-adopter/)

De nombreux éléments entrant en compte dans la parité en entreprise n'ont pas été pris en compte dans cette modélisation et pourraient être ajoutés. Nous ne les avons pas mis en place par manque de temps, mais aussi à cause des difficultés de modélisation et du manque de chiffres sur lesquels nous appuyer.

### La rédaction de l'annonce
Les annonces doivent être rédigées de façon neutre : pas de masculin par défaut, éviter les adjectifs associés à des clichés de genre, etc. La loi impose notamment la mention "F-H ou H-F" dans les offres d'emploi[<sup>test</sup>](https://analyseur.acompetenceegale.com/comment-eviter-discriminations-a-lembauche-selon-sexe/).
Nous aurions pu modifier la proportion de femmes ou d'hommes postulant pour une offre en fonction de la formulation de l'annonce.

### La présentation de l'entreprise
Une entreprise devrait mettre autant en avant des employés hommes que des employées femmes sur les sites de présentation de l'entreprise, et ceux dans tous les domaines d'activité (c'est-à-dire, ne pas représenter des femmes que pour les postes en ressources humaines, ou que des hommes pour les postes considérés comme plus techniques).
Nous aurions pu modifier la proportion de femmes ou d'hommes postulant pour une offre d'emploi en fonction de la proportion de femmes et d'hommes représenté.e.s sur le site de l'entreprise.

### L'anonymisation des candidatures
Dans un premier temps du processus de recrutement de l'entreprise, il est conseillé d'anonymiser les candidatures.
Nous aurions pu modifier la proportion de femmes ou d'hommes convoqué.e.s en entretien en fonction de si l'entreprise l'anonymise les candidatures.

### Les entretiens
Avoir une équipe avec autant de femmes que d'hommes, à la fois pour limiter les biais lors du choix de recrutement, mais aussi pour montrer une plus grande diversité aux candidat.e.s participant au processus de recrutement. De plus, on peut penser à prendre en entretien autant de femmes que d'hommes.
Nous aurions pu modifier qui reçoit une offre d'emploi sur ce critère, ou encore, modifier la réponse du/de la candidat.e sélectionné.e.

### Les avantages au sein de l'entreprise
L'entreprise peut proposer des avantages comme une crèche au sein de l'entreprise, des horaires flexibles, du télétravail, etc.
Ce sont des critères pris en compte par les candidats au moment de postuler et d'accepter une offre d'emploi, que nous aurions pu implémenter.

### Des mesures anti-VSS
Les entreprises peuvent prendre des mesures contre les violences sexistes et sexuelles. Cela peut aller de la distribution de prospectus, à des formations sur le sujet et la mise en place de cellules dédiées.
Nous aurions pu prendre ces éléments en considération sur la façon dont les agressions sont gérées (sanctions différentes, etc.), et éventuellement modifier les probabilités qu'une agression ait lieu (une personne ayant suivi une formation a moins de chance d'agresser, etc.).

### Les VSS
Notre modélisation s'appuie sur des chiffres concernant les agressions sexuelles **A FAIRE vérifier ça**. Nous aurions pu prendre en compte toutes les VSS et changer l'impact sur la santé mentale en fonction des différents types de VSS (et aussi changer les sanctions pour l'employé.e qui les a commises).

### L'intervention du/de la psychologue d'entreprise
Lors de signalement pour violence sexiste ou sexuelle, le personne ayant déposée le signalement a le droit à un accompagnement par la.e psychologue de l'entreprise.
Nous aurions pu modéliser à quel point cet accompagnement est utilé avec une hausse de santé mentale.

### Les causes de départs
Nous prenons en compte les départs après les congés maternités, mais nous n'avons pas de chiffres pour les hommes. Nous ne prenons pas en compte les congés sans solde (la personne est toujours dans l'entreprise, mais ne travaille pas et ne perçoit pas de salaire), ni toutes les causes de départ. En particulier, il aurait été intéressant de prendre en compte les départs des employé.e.s s'occupant de proches malades (on suppose que les femmes partent plus souvent que les hommes dans ce cas).

### Les différences de salaire
Dans notre modélisation, tous les employé.e.s ont le même salaire.
Pour être au plus proche de la réalité, il aurait fallu prendre en compte les différents postes, les augmentations et promotions.

### Le secteur
L'entreprise modélisée n'a pas de secteur dédié : tous les chiffres utilisés sont des chiffres généraux, or, ils varient fortement d'un secteur à l'autre.
Nous aurions pu laisser le choix à l'utilisateur.ice du secteur souhaité et prendre en compte les chiffres correspondant.

### Pourquoi nous n'avons pas utilisé l'index de l’égalité professionnelle entre les femmes et les hommes ?
Cet index mis en place par le gouvernement, et devant être partagé tous les ans par les entreprises de plus de 50 salarié.e.s, permet de calculer l'égalité professionnelle entre les femmes et les hommes dans une entreprise. Il repose sur cinq indicateurs[<sup>test</sup>](https://travail-emploi.gouv.fr/droit-du-travail/egalite-professionnelle-discrimination-et-harcelement/indexegapro) :
>   * L’écart de rémunération femmes-hommes,
>   * L’écart de répartition des augmentations individuelles,
>   * L’écart de répartition des promotions (uniquement dans les entreprises de plus de 250 salariés),
>   * Le nombre de salariées augmentées à leur retour de congé de maternité,
>   * La parité parmi les 10 plus hautes rémunérations.
Cet index repose donc principalement sur des salaires et des promotions et augmentations que nous n'avons pas modélisées.

## Points à améliorer sur la simulation actuelle
En plus des ajouts possibles mentionnés dans la partie précédentes, des points actuelles de la simulation peuvent être améliorés.

### Sur la modélisation
* Le bénéfice : nous avons pris des chiffres très généraux sur les coûts des salarié.e.s, du recrutement et les bénéfices générés par les employé.e.s ;
* La montée de productivé liée à la présence d'hommes : nous n'avons pas de chiffres sur l’intérêt d’avoir des hommes sur la bonne ambiance dans l'entreprise et ne l'avons donc pas modélisé ;
* Les départs après un congé paternité : nous n'avons un chiffre que pour les départs après un congé maternité ;
* Le teambuilding : on modélise boost positif pour les employé.e.s lors de l'organisation d'un teambuilding, mais nous n'avons pas de chiffre pour appuyer cette modélisation, et toutes les entreprises ne font pas de teambuilding ; 
* Le recrutement : nous engageons chaque année 5% d'employé.e.s supplémentaires, mais c'est un chiffre décidé arbitrairement, de plus, nous considérons que les postes seront toujours pourvus, et nous ne cherchons pas à remplacer les personnes qui ont quitté l'entreprise (le recrutement est fait indépendamment des départs et les embauches représentent toujours une hausse de 5% de l'effectif total) ;
* L'amende liée à l'absence de femme : il s'agit d'une amende liée à la loi de Rixain qui est prise en compte dans notre modélisation, mais cette loi ne s’appliquera qu’à partir de 2026 ;
* Les méthodes de recrutement : pour rappel, les places réservées n'existent pas dans la réalité ;

### Sur la simulation
* Les graphes : au départ, nous voulions un graphe avec toutes les données ensemble, mais il a été supprimé à cause d'un problème d'échelle (elle n'était pas la même pour tous les graphes), mais nous pourrions trouver une solution à ce problème.
