#!/usr/bin/env bash

# include my library helpers for colorized echo and require_brew, etc
source ./lib/echobot.sh

function doDeploy(){
	case "$1" in
	    all )
	        doDeploy bridge
	        doDeploy center
	        doDeploy farm
	        doDeploy hyper
	        doDeploy nuke
	        doDeploy town
	    	;;
	    bridge )
	    	bot "deploying bridge"
	        scp -r drawbridge.club/* antic@drawbridge.club:~/
	        #scp -r shared/public_html/* antic@drawbridge.club:~/public_html/
	        ok ;;
	    center )
	    	bot "deploying center"
	        scp -r tilde.center/* antic@tilde.center:~/
	        ok ;;
	    farm )
	    	bot "deploying farm"
	        scp -r tilde.farm/public_html/* antic@tilde.farm:~/public_html/ 
	        #scp -r shared/public_html/* antic@tilde.farm:~/public_html/
	        ok ;;
	    hyper )
	    	bot "deploying hyper"
	        scp -r hypertext.website/* antic@hypertext.website:~/ 
	        #scp -r shared/public_html/* antic@hypertext.website:~/public_html/
	        ok ;;
	    nuke )
	    	bot "deploying nuke"
	        scp -r totallynuclear.club/* antic@totallynuclear.club:~/ 
	        #scp -r shared/public_html/* antic@totallynuclear.club:~/public_html/
	        ok ;;
	    town )
	    	bot "deploying town"
	        scp -r tilde.town/* antic@tilde.town:~/ 
	        #scp -r shared/public_html/* antic@tilde.town:~/public_html/
	        ok ;;
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