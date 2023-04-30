# Objective

Data validation for entering the system.

In the controller, the types that live here are used to talk to the outside world like endpoint
requests.

(I'm not sure about this now) They could be used to talk to the database
which is also the outside world.

But in the MVC architecture the database or persistence layer could be different
compared to the hexagonal architecture.

I understand (I should confirm this later) that the Repository is the package responsible for hosting the code that is responsible for interacting with the persistence layer, that is, database.
