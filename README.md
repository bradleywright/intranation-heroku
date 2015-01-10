# intranation.com on Heroku

My blog is currently mostly hosted on Tumblr. All my own server was running was
an Nginx server that redirected all of my old links so they didn't change or
break links.

Given that, it makes more sense to host the redirection for free on Heroku and
shut down my server. Hence this Go app.

Built using
[the 'official' guide](http://mmcgrana.github.io/2012/09/getting-started-with-go-on-heroku.html).
