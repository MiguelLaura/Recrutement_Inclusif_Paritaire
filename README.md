**ATTENTION REMPLACER LES TEST PAR NUMERO + FAIRE LES A FAIRE**

# La parité en entreprise

## Le contexte de réalisation du projet
Ce projet a été réalisé dans le cadre de l'UV IA04 enseignée à l'Université de Technologie de Compiègne. Il a été réalisé par :
* Mathilde Lange
* Solenn Lenoir
* Nathan Menny
* Laura Miguel

## Le sujet
Ce projet est une simulation d'une entreprise. On souhaite étudier la parité dans une entreprise. Nous nous sommes en particulier concentré.e.s sur le recrutement : c'est une étape systématique dans une entreprise qui peut influer sur la parité. Le but est donc de laisser l'utilisateur.ice choisir une façon d'effectuer son recrutement, de lancer la simulation et voir l'influence de ce recrutement sur le bénéfice de l'entreprise, le nombre de départs, etc.

## La problématique
Notre principale question est donc : quelles sont les conséquences d’un recrutement tenant compte de la parité sur la santé d’une entreprise ?

## Récupérer le projet et le lancer
**A FAIRE compléter**
```bash
go get github.com/gorilla/websocket
go run cmd/launch-all/launch-all.go
```

## La table des matières

* [Le contexte de réalisation du projet](#le-contexte-de-réalisation-du-projet)
* [Le sujet](#le-sujet)
* [La problématique](#la-problématique)
* [Récupérer le projet et le lancer](#récupérer-le-projet-et-le-lancer)
* [La table des matières](#la-table-des-matières)
* [L'interface de simulation](#linterface-de-simulation)
    * [Le formulaire](#le-formulaire)
        * [Pourquoi un pourcentage de femmes à atteindre ?](#pourquoi-un-pourcentage-de-femmes-à-atteindre)
        * [Places réservées ?](#places-réservées)
        * [Compétences égales ?](#compétences-égales)
    * [La simulation](#la-simulation)
    * [Fonctionnement de l'interface](#fonctionnement-de-linterface)
* [La modélisation](#la-modélisation)
    * [Ce qui est modélisé et les sources](#ce-qui-est-modélisé-et-les-sources)
        * [Employé](#employé)
        * [Recrutement](#recrutement)
        * [Agressions, plaintes et licenciement](#agressions-plaintes-et-licenciement)
        * [Départs (hors licenciement)](#départs-hors-licenciement)
        * [Congés maternité](#congés-maternité)
        * [Formation](#formation)
        * [Team building](#team-building)
        * [Bénéfices](#bénéfices)
    * [L'exprimer dans le code](#lexprimer-dans-le-code)
* [Les résultats](#les-résultats)
* [Non pris en compte dans notre modélisation](#non-pris-en-compte-dans-notre-modélisationtest)
    * [La rédaction de l'annonce](#la-rédaction-de-lannonce)
    * [La présentation de l'entreprise](#la-présentation-de-lentreprise)
    * [L'anonymisation des candidatures](#lanonymisation-des-candidatures)
    * [Les entretiens](#les-entretiens)
    * [Les avantages au sein de l'entreprise](#les-avantages-au-sein-de-lentreprise)
    * [Les mesures anti-VSS](#les-mesures-anti-vss)
    * [Les VSS](#les-vss)
    * [L'intervention du/de la psychologue d'entreprise](#lintervention-dude-la-psychologue-dentreprise)
    * [Les causes de départs](#les-causes-de-départs)
    * [Les différences de salaire](#les-différences-de-salaire)
    * [Le secteur](#le-secteur)
    * [Pourquoi nous n'avons pas utilisé l'index de l’égalité professionnelle entre les femmes et les hommes ?](#pourquoi-nous-navons-pas-utilisé-lindex-de-légalité-professionnelle-entre-les-femmes-et-les-hommes)
* [Remarques sur la simulation actuelle](#remarques-sur-la-simulation-actuelle)
    * [Sur l'interface](#sur-linterface)
    * [Sur la modélisation](#sur-la-modélisation)

## L'interface de simulation

### Le formulaire
**A FAIRE mettre une capture ?**
Nous avons en premier temps un formulaire dans lequel l'utilisateur.ice entre les informations mentionnées ci-dessous. Il y a également des informations sur la simulation qui correspondent à la partie [Le sujet](#le-sujet) et cette partie.

L'utilisateur.ice va pouvoir définir :
* le nombre d'employé.e.s de l'entreprise,
* la durée de la simulation (un pas de la simulation correspond à une année dans l'entreprise),
* le pourcentage initial de femmes dans l'entreprise,
* s'iel a un pourcentage de femmes à atteindre,
* le type de recrutement (s'iel a certaine répartition femmes-hommes à atteindre, il faut choisir un type de recrutement pour quand on est en dessous de ce pourcentage, et un type de recrutement pour quand on est au-dessus du pourcentage).

#### Pourquoi un pourcentage de femmes à atteindre ?
Les entreprises peuvent vouloir atteindre une certaine répartition femmes-hommes (pour respecter une loi, favoriser l’innovation, etc.) et mettre en place des stratégies temporaires, notamment au niveau du recrutement. Ainsi, les manières de recruter ne seront pas les mêmes en dessous ou au-dessus du seuil défini.

#### Places réservées ?
Pour cette stratégie, sur le nombre de personnes à recruter, on choisira de recruter un pourcentage fixe de femmes ou d'hommes (on prendra toujours les plus compétent.e.s dans cette population) puis, pour le reste des candidat.e.s, on recrutera en fonction des compétences seulement. Cette stratégie n'existe pas dans la réalité puisqu'il s'agit d'une discrimination de genre. En effet, on ne peut discriminer à l'embauche sur le genre que pour des cas particuliers, comme pour le cinéma ou mannequinat[<sup>test</sup>](https://analyseur.acompetenceegale.com/comment-eviter-discriminations-a-lembauche-selon-sexe/).

#### Compétences égales ?
Pour cette stratégie, on recrute d’abord la personne la plus compétente. Si jamais deux personnes ont des compétences équivalentes, on choisira qui recruter en fonction de ce qui a été demandé par l’utilisateur.ice : iel choisit s'iel donne sa préférence à une femme, à un homme, ou s'iel n'a pas de préférence et prend un.e des candidat.e.s au hasard. C'est un type de recrutement qu'on peut appliquer à la vie réelle, mais uniquement en cas de candidatures comparables, en faveur du genre sous-représenté et en cas de dernier critère de départage[<sup>test</sup>](https://egaliteautravail.com/domaine/recrutement/).

### La simulation
**A FAIRE mettre une capture ?**
La validation du formulaire nous renvoie sur la page de simulation. Nous pouvons alors la lancer (soit de façon à ce que les pas s'enchaînent sans action de l'utilisateur.ice, soit en avançant pas à pas). On peut également arrêter la simulation, la mettre en pause et revenir au formulaire.
Quand la simulation est lancée, on peut voir depuis combien d'années l'entreprise tourne sous la simulation, le nombre d'employé.e.s, la parité, les bénéfices. En particulier, on a des graphes nous montrant l'évolution, au cours des années, des bénéfices, de la parité, des compétences des employé.e.s et de la santé mentale des employé.e.s.
Dans une partie *Tableau de bord*, on peut voir des informations sur ce qu'il se passe au cours des années. Ces informations sont divisées en catégories (on peut sélectionner les catégories qu'on souhaite voir dans le tableau de bord) :
* Agression : le nombre d'agressions entre employé.e.s et le nombre de signalements faits auprès de l'entreprise ;
* Départ : le nombre de démissions (spontanées, dûe à une dépression ou après un congé maternité), le nombre de retraites, le nombre de licenciements ;
* Entreprise : si l'entreprise reçoit des amendes liées à sa parité, ou un bonus de productivité ;
* Recrutement : le nombre d'embauches et des détails sur le comportement des ressources humaines pendant le processus de recrutement ;
* Employé : le nombre de naissances d'enfants et de congés parentaux ;
* Evénements : l'organisation de team building et le nombre d'employé.e.s ayant participé à une formation.
On a également des informations au survol sur le bénéfice, le recrutement et les catégories pour avoir des explications supplémentaires.

### Fonctionnement de l'interface

L'interface a été réalisée en HTML/CSS/JavaScript. Nous utilisons la bibliothèque Chart.js pour créer les graphes et visualiser les données au cours du temps. Pour envoyer les informations issues du formulaire et créer une nouvelle simulation, nous utilisons une requête POST. Les informations envoyées et retournées peuvent être consultées dnas le document [API.md](API.md).

Une fois sur la page de la simulation, toutes les informations sont transférées grâce à des websockets. Les données sont de différents types et envoyées à différents moments. Nous utilisons un Logger qui envoie les données dans les websockets, en même temps qu'il les affiche dans la console. Ce Logger est commun à la simulation et à tous les agents (employé-es, recrutement et entreprise).

D'abord, les informations "initiales" qui concernent la simulation créée (l'id de la simulation, le nombre d'années, le type de recrutement choisi, status de la simulation, ...) sont envoyées à la page HTML dès qu'une connexion websocket est établie. Cela permet à la simulation de n'être pas dépendante d'une page web particulière. Ainsi, on peut toujours retrouver les informations lorsqu'on reload une page web avec l'id de la simulation. Dans le code, ces informations sont regroupés dans "LOG_INITAL".

Nous gérons également des informations sur le status de la simulation, par exemple si elle est terminée, à relancer, en pause, pour afficher ces informations sur l'interface avec des popup temporaires. Ce sont des "LOG_REPONSE". 

Ensuite, pour afficher les informations au fur et à mesure, la simulation envoie chaque année son pas de temps actuel, son nombre d'employé-es, son pourcentage de femmes, son bénéfice, la moyenne des compétences et la moyenne de la santé mentale. Cela correspond aux informations globales, "LOG_GLOBAL". 

Enfin, pour suivre le déroulé de la simulation, chaque agent (que ce soit des employé-es, le recrutement ou l'entreprise) log des informations lorsqu'il agit. Cela permet de garder les informations sur les actions au moment où elles sont réalisées. Ces logs sont catégorisés pour pouvoir être affichés de différentes couleurs dans l'interface et masqués si besoin. L'entreprise peut envoyer les logs suivants : LOG_AGRESSION, LOG_DEPART, LOG_ENTREPRISE, LOG_EVEMENT. Le recrutement envoie les LOG_RECRUTEMENT et les employé-es LOG_EMPLOYE.

## La modélisation

### Ce qui est modélisé et les sources
**A FAIRE expliquer les sources et ce qu'on a modélisé**

#### Employé
Les compétences d'un employé sont modélisées par un entier 0 à 10. Elles suivent une loi normale avec mu=5 et sigma=3. Cela permet d'obtenir une majorité d'individus moyens et assez peu d'individus excellents ou mauvais[<sup>test</sup>](https://www.ruf.rice.edu/~lane/papers/male_female.pdf).
Pour la santé mentale, nous avons décidé de la modéliser comme un entier de 0 à 100. Lorsqu'il rejoint l'entreprise, il dispose d'une santé mentale pleine de 100. Son expérience au sein de l'entreprise peut augmenter ou réduire sa santé mentale dans le domaine défini.
Concernant l'ancienneté, il s'agit d'un entier entre 0 et 40. Les employés générés au début de la modélisation ont une ancienneté aléatoire. Les nouveaux employés ont une ancienneté de 0. Cette valeur est incrémentée tous les ans. Lorsqu'un employé a une ancienneté de 40, il part à la retraite.
**CHANGER VALEUR ANCIENNETE MAX, lire des trucs sur la retraite moyenne**

https://www.harcelement.eu/les-statistiques-choquantes-sur-le-harcelement-sexuel-au-travail-en-france/

#### Recrutement

#### Agressions, plaintes et licenciement
Chaque agresseur agresse tous les ans, uniquement une fois, quelqu'un qui n'est pas lui. Un homme agressera une femme et vice versa sauf si ne trouve pas donc agressera même sexe. Si ne trouve personne à agresser à part lieu même 4 fois, il n'agresse pas.
Une personne agressé 

#### Départs (hors licenciement)

#### Congés maternité

#### Formation

#### Team building

#### Bénéfices

### L'exprimer dans le code
**A FAIRE insérer diagramme de classe et séquence, parler des loggers**

#### Simulation
Cette classe gère la simulation et en particulier le lien entre le front et la back.

**A FAIRE parler en particulier de la gestion de la communication**

#### Entreprise
L'entreprise est un agent qui assume aussi le rôle de l'environnement puisqu'elle gère les différents agents et c'est elle qui centralise les informations.

Lorsqu'on lance l'entreprise, elle lance les employés et le recrutement, et elle entre dans une boucle. Au début de la boucle, l'entreprise attend un message sur le channel le liant à la simulation. Le message est soit `LIBRE` et ce qui indique à l'entreprise de lancer une année (ce qui revient à lancer la méthode `agir`), soit `FIN` et ce qui lui indique d'arrêter tous les agents qu'elle a lancé. Si le nombre d'employés atteint zéro avant la fin de la simulation, on passe tous les step sans rien faire jusqu'à la fin. Quand l'entreprise sort de la boucle, elle attend un message final de la simulation pour confirmer l'arrêt.

Quand elle lance une année, l'entreprise :
* organise les formations ;
* organise un team building ;
* lance les recrutements en envoyant un message à l'agent recrutement avec le nombre d'employés à recruter (le recrutement a un channel dédié à la communication entreprise->recrutement) ;
* envoie un message aux employés pour leur dire d'agir (chaque employé a un channel dédié à la communication entreprise->employé).
L'entreprise attend un message de chacun des employés (entreprise a un channel dédié aux retours des employés) pour lancer la fin d'année. Pour attendre les employés, l'entreprise a une fonction `RecevoirActions` qui prend le nombre d'actions à recevoir en entrée, lance une boucle dans laquelle elle attend un message sur son channel et incrémente un compteur à chaque fois qu'elle reçoit un message. Quand le compteur atteint le nombre d'actions passé en entrée, c'est la fin de la boucle.
En fin d'année, l'entreprise lance un team building, et lance la fin d'année, elle gère :
* les plaintes (donc licencie éventuellement des employés) ;
* les départs ;
* les recrutements : elle attend un message du recrutement qui l'informe de la fin du processus de recrutement (entreprise a un channel dédié aux retours du recrutement), puis lance les nouveaux employés.

Pour arrêter tous les agents, l'entreprise envoie un message de fin sur les channels des employés et sur le channel du recrutement. Elle doit attendre leurs retours avant de s'arrêter elle-même.

Au cours des actions des employés, ceux-ci peuvent changer les listes des employés démisionaires, la liste des départs, la liste des plaintes, le nombre de dépressions et le nombre de congés parentaux : pour éviter des problèmes d'accès concurents, les fonctions gérant ces changements posent un `Lock` sur l'entreprise.

#### Employé
Les employés sont des agents lancés par l'entreprise.

Quand ils sont lancés, ils entrent dans une boucle dans laquelle ils effectuent leurs actions jusqu'à ce qu'ils aient été arrêtés par l'entreprise.
Pour agir, les employés attendent un message de l'entreprise sur un channel dédié. Si l'entreprise leur envoie le message `LIBRE`, ils vont :
* agresser s'ils sont agresseurs ;
* se former s'ils sont dans la liste des employés recevant une formation pendant l'année en cours ;
* vieillir ;
* potentiellement avoir un enfant et donc potentiellement démissionner après un congé maternité ;
* partir à la retraite s'ils ont assez d'ancienneté ;
* potentiellement poser une démission spontannée.
À chaque action, ils vont modifier des informations centralisées par l'entreprise, donc pour se faire, ils appellent des méthodes de l'entreprise.
Si le message reçu par l'employé de la part de l'entreprise est `AGRESSION`, cela signifie que l'employé se fait agresser. Il va donc perdre de la santé mentale et potentiellement porter plainte auprès de l'entreprise et partir en depression.
Si le message est `FIN`, l'employé passe son attribut fin à `true` ce qui lui permettra de sortir de la boucle de vie.
Enfin, l'employé envoie un message à l'entreprise sur un channel pour l'informer qu'il a fini ses actions.

#### Recrutement
Le recrutement est un agent crée par la simulation et lancé par l'entreprise.

**A FAIRE**

## Les résultats
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

### Les mesures anti-VSS
Les entreprises peuvent prendre des mesures contre les violences sexistes et sexuelles. Cela peut aller de la distribution de prospectus, à des formations sur le sujet et la mise en place de cellules dédiées.
Nous aurions pu prendre ces éléments en considération sur la façon dont les agressions sont gérées (sanctions différentes, etc.), et éventuellement modifier les probabilités qu'une agression ait lieu (une personne ayant suivi une formation a moins de chance d'agresser, etc.).

### Les VSS
Notre modélisation s'appuie sur des chiffres concernant les agressions sexuelles **A FAIRE vérifier ça**. Nous aurions pu prendre en compte toutes les VSS et changer l'impact sur la santé mentale en fonction des différents types de VSS (et aussi changer les sanctions pour l'employé.e qui les a commises).

### L'intervention du/de la psychologue d'entreprise
Lors de signalement pour violence sexiste ou sexuelle, le personne ayant déposée le signalement a le droit à un accompagnement par la.e psychologue de l'entreprise.
Nous aurions pu modéliser à quel point cet accompagnement est utile avec une hausse de santé mentale.

### Les causes de départs
Nous prenons en compte les départs après les congés maternités, mais nous n'avons pas de chiffres pour les hommes. Nous ne prenons pas en compte les congés sans solde (la personne est toujours dans l'entreprise, mais ne travaille pas et ne perçoit pas de salaire), ni toutes les causes de départ. En particulier, il aurait été intéressant de prendre en compte les départs des employé.e.s s'occupant de proches malades (on suppose que les femmes partent plus souvent que les hommes dans ce cas).

### Les différences de salaire
Dans notre modélisation, tous les employé.e.s ont le même salaire.
Pour être au plus proche de la réalité, il aurait fallu prendre en compte les différents postes, les augmentations et promotions.

### La hiérarchie des postes
Nous n'avons pas modélisé de hiérarchie de postes. Or, on pourrait supposer que si une entreprise a plus de femmes à la direction, et que celles-ci sont intéressées pour embaucher des femmes, elles auraient plus de pouvoirs faviriser leur recrutement. Ou encore que dans le cas de VSS commises par des haut placés, celles-ci sont moins signalées. C'est donc un point qu'il aurait été intéressant d'étudier.

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

## Remarques sur la simulation actuelle
En plus des ajouts possibles mentionnés dans la partie précédentes, des points actuelles de la simulation peuvent être améliorés.

### Sur l'interface
* Les graphes : au départ, nous voulions un graphe avec toutes les données ensemble, mais il a été supprimé à cause d'un problème d'échelle (elle n'était pas la même pour tous les graphes), mais nous pourrions trouver une solution à ce problème.

### Sur la modélisation
* Le bénéfice : nous avons pris des chiffres très généraux sur les coûts des salarié.e.s, du recrutement et les bénéfices générés par les employé.e.s ;
* La montée de productivé liée à la présence d'hommes : nous n'avons pas de chiffres sur l’intérêt d’avoir des hommes sur la bonne ambiance dans l'entreprise et ne l'avons donc pas modélisé ;
* Les départs après un congé paternité : nous n'avons un chiffre que pour les départs après un congé maternité ;
* Le team building : on modélise boost positif pour tous les employé.e.s lors de l'organisation d'un team building (ce qui n'est pas forcément le cas dans la réalité), mais nous n'avons pas de chiffre pour appuyer cette modélisation, et toutes les entreprises ne font pas de team building ;
* Le recrutement : nous engageons chaque année 5% d'employé.e.s supplémentaires, mais c'est un chiffre décidé arbitrairement, de plus, nous considérons que les postes seront toujours pourvus, et nous ne cherchons pas à remplacer les personnes qui ont quitté l'entreprise (le recrutement est fait indépendamment des départs et les embauches représentent toujours une hausse de 5% de l'effectif total) ;
* L'amende liée à l'absence de femme : il s'agit d'une amende liée à la loi de Rixain qui est prise en compte dans notre modélisation, mais cette loi ne s’appliquera qu’à partir de 2026 ;
* Les méthodes de recrutement : pour rappel, les places réservées n'existent pas dans la réalité.
