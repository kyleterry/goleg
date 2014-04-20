#Goleg - Go bindings for liboleg

Goleg is a library/wrapper/bindings/wrabings/bindpers of liboleg/OlegDB.

You can either use the direct lowlevel bindings (wrapper.go) as they have very similar apis to the [original library itself](https://olegdb.org/docs/0.1.0/en/documentation.html#functions) or use the more convenient high-level ones (highlevel.go).

##FAQ

###What the hell is OlegDB?
[OlegDB](https://olegdb.org) is a k/v store database that tries to be awful, but the programmers are terrible at writing awful code, so it's turning out quite well.  
It's written in C89 + POSIX, so it will probably compile on anything decent without [50 thousand lines of code of compatibility hacks](http://opensslrampage.org/).. so there's that.

###Who would ever use this anyway?
I don't know! But I've written a [Go frontend](https://github.com/hamcha/golegd) that can replace OlegDB's builtin one.. and is a little bit faster.
