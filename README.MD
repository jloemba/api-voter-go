
# Projet Go
##Installation

Utilisez le makefile pour:
- initialiser le projet : $ make build
- lancer le projet : $ make run
- supprimer le docker : $ make down

##Postman
#####pour une utilisation postman voici les différentes routes avec des examples de json:
<br>

Post: localhost:8081/api/user/new

{
"first_name":"Bruce",
"last_name":"Wayne",
"email":"t@wayne.ioa",
"password":"marthajjjj88",
"birthdate":"1995-01-01T00:00:00Z"
}

<br>

Post: localhost:8081/api/user/login

{
"email":"t@wayne.ioa",
"password":"marthajjjj88"
}

<br>

Delete: localhost:8081/api/user/delete/{uuid_voulu}

<br>

Put: localhost:8081/api/user/put/{uuid_voulu}

{
"first_name":"Brucee",
"last_name":"Waynee",
"email": "t@wayne.io",
"password":"marthajjjj77777"
}

<br>

Post: localhost:8081/api/vote/create

{
"title":"Propreté des trottoirs",
"description":"Dans le budget qui sera soumis au vote des conseillers de Paris lundi, 32 M€ seront consacrés à l’achat de nouvelles machines, plus silencieuses, plus propres et plus performantes. Il n’y aura pas d’embauche d’agents supplémentair.",
"start_date":"2019-12-01T17:23:05.636849176+01:00",
"end_date":"2019-12-10T17:23:05.636849176+01:00"
}

<br>

Delete: localhost:8081/api/vote/delete/{id_voulu}

<br>

Put: localhost:8081/api/vote/update/{id_voulu}

{
"title":"Présidentiel United States",
"description":"Campagne présidentiel United States"
}

<br>

Get: localhost:8081/api/vote/show/{uuid}

