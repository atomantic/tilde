#!/usr/bin/env bash

# include my library helpers for colorized echo and require_brew, etc
source ./lib/echobot.sh

function push(){
	bot "deploying to "$1
	rsync -avz --exclude '.DS_Store' ./~/$1/* antic@$1:~/
	rsync -avz --exclude '.DS_Store' --delete shared/public_html/* antic@$1:~/public_html/
	ok
}

function doDeploy(){
	case "$1" in
	    all )
	        doDeploy bridge
	        doDeploy center
	        doDeploy cybyte
	        doDeploy farm
	        doDeploy german
	        doDeploy hyper
	        doDeploy nuke
            doDeploy retro
	        doDeploy town
	        doDeploy s3
	    	;;
	    bridge )
	    	push drawbridge.club
	        ;;
	    cybyte )
	        push cybyte.club
	        ;;
	    center )
	        push tilde.center
	        ;;
	    farm )
	        push tilde.farm
	        ;;
	    geman )
			push germantil.de
			;;
	    hyper )
	        push hypertext.website
	        ;;
	    nuke )
	        push totallynuclear.club
	        ;;
        retro )
            push retronet.net
            ;;
	    s3 )
	        aws s3 cp ./~/s3/* s3://tildetowneivy
	        aws s3 cp ./shared/public_html/ s3://tildetowneivy/ --recursive
	        ;;
	    town )
	        push tilde.town
	        ;;
	    *)
	    	bot "you can run 'deploy' with the following options:\n"
	    	echo "bridge, hyper, farm, town, nuke, all"
	    	;;
	esac
}

case "$1" in
    deploy)
        doDeploy $2
        ;;
    *)
        bot "Hi, I can go do some things.\n"
        echo -e "Run each as an argument to this script (e.g. 'go deploy nuke'):\n"
        echo "deploy {dest} - deploy to a destination tilde"
        ;;
esac