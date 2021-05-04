This is the code related to the project/blog at <http://ericski.com/tags/spellerproject/>


In software engineering, it is very common to have services that are backed by some sort of datastore. Typically these are
some sort of database whether it is a traditional RDBMS or a newer NOSQL solution. However, if you have static data, plain old
files can be a perfectly good substitute.

For this sample project, we are going to write a simple spelling validation application. We will primarily be focused on the backend services
because this doesn\'t really need too much in the way of UI and, if we\'re being honest, my UI skills are weak at best. 😂

The service is quite simple, consisting of a single GET endpoint:

- /{word} returns true if {word} is a real word, false otherwise

Eventually, we will need to add some other endpoints because we are going to be making this more of a production quality service
that runs in kubernetes.

As for programing language choice, we will be using Go because it has a couple "cloud friendly" qualities.

- small footprint - memory and cpu
- single file delivery - easy to deploy/package
- good support for writing REST services
