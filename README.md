# all my ~

- [cybyte.club/~antic](http://cybyte.club/~antic)
- [drawbridge.club/~antic](http://drawbridge.club/~antic)
- [germantil.de/~antic](http://germantil.de/~antic)
- [hypertext.website/~antic](http://hypertext.website/~antic)
- [retronet.net/~antic](http://retronet.net/~antic)
- [tilde.center/~antic](http://tilde.center/~antic)
- [tilde.town/~antic](http://tilde.town/~antic)
- [tilde.farm/~antic](http://tilde.farm/~antic)
- [totallynuclear.club/~antic](http://totallynuclear.club/~antic)

I'd be delighted to join your tilde server. My username is ~antic and you can get [my public ssh keys are here](https://raw.githubusercontent.com/atomantic/tilde/master/shared/.ssh/authorized_keys)

You can find me on the mainstream web in many other places:

- [portfilio v2](https://atomantic.github.io)
- [old portfolio](http://adameivy.com)
- [twitter: @antic](https://twitter.com/antic)
- [facebook: @atomantic](https://twitter.com/atomantic)


## Notes on this project

- I'm currently converting this project to be more DRY and more generically useful for managing multiple ~ accounts
- Go ahead and fork and pull-request if you like :)

## TODO

- generate `go` bash script from a `gulp setup` command
	- read tilde sites from ~.json and allow commands configured via that
- pull out all copies of similar content to shared and template directories
- generate shared/allmy~.html from ~.json
- add gulp build task to generate ~ index.html sites using shared template content

# Structure


# Commands
Assuming `./` is in your PATH (otherwise, you will have to `./go`)

Deploy all ~ sites:
```
go deploy all
```

Deploy a single ~ site:
```
go deploy nuke
```