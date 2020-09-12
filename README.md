# Radio (Event-Bus for Golang)

Radio is a small framework for managing an event base system on application level. You can easily notify another component
of you application. It allows you to fire an event from one part of your application and spread it to different parts.

Please feel free to enhance the framework by contributing. (Just fork it and make a PR!)

## Basic Concepts of the Framework

## Concept

Radio does provide you a possibility to dispatch an event from one side of you application to all corners. This allows
you to react on events happened without having dependencies between your internal application components. If you want to
know more in this topic you could have a look in to the topic of
[Event-driven programming](https://en.wikipedia.org/wiki/Event-driven_programming).

By creating a `Radio` object you can attach new channels with a path to it. The structure hiding behind it is a tree. 
Which means that one `Channel` can have multiple children and may has a parent channel, and this channel then do may also
have another time multiple more channels.

This means also for `Event`s, that it is passed into the channel you wanted plus all children channels. Events passed to
the child of a parent, can never be pushed to it's parent (or only if you didn't do it you self).

You also do can have multiple of the radios in you application. Normally it does make sense to instead setup multiple
*main* channels and have one child for each of your bigger component. If you then want to split you component into
multiple fragments (which still need to communicate together), you can do it the same way for you channels, by just
adding more child to the component channel.
