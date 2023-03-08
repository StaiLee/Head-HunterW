# README

## About

Groupie tracker  StaiLee

## how to run ?

need :
    Go 1.18+
    NPM (Node 15+)

then :
    go go install github.com/wailsapp/wails/v2/cmd/wails@latest
    wails doctor ( pour verifier si tout est bien installer)


## command

voici les commandes wails : 
    -wails dev : pour "run" l'application
    -wails build : pour compiler le projet et obtenir l'executable...

## documentation

pour plus d'information le lien de la doc :

https://wails.io/fr/

## fonctionnement

l'application bureau n'est enfaite que le reflet d'un serveur web en golang, nous allons donc a l'aide des templates html que nous avons l'habitude d'utiliser en golang
parse les donnees de l'api fournis, c'est dernier vont afficher a l'aide de notre app.svelte, vous pouvez donc retrouver le serveur au port 1706 de votre navigateur avec 
les differents chemins pour plus de details, wails nous permet de fournir un travail en GUI, mais le fonctionnement reste comme si nous travaillons sur du web

vous avez a disposition differentes fleches en haut pour pouvoir naviguer dans l'app il ya un index avec tout les artistes et les donnees qui vont avec et une template
avec la barre de recherche pour trier et afficher seulement l'information desirer
