# relationship invariants
Difficult

Index all instances on their primary keys during validation.
Thus an association can lookup the instance and allow navitation to it and its properties.
Very handy.

However...
The instance is not required to be in the document being validated. It is simply assumed that when storing the
instance and its relationship that the instance it is referencing exissts in the db. That is why a driver
first adds all instances and then processes the relationships (as they would fail if pointing to non
existing data.

Thus in order to fully validate the validator must have access to a provider of information from
the DB.

Options:
* Do the full DB thing
* Do not provide cross object validation for associations (except the edge properties)
* Does it make sense to only validate invariants in the document? [No, not really]
* Rely on triggers in the DB? [does not really work well]


"A car must have 4 wheels" - if the document has 4 it would validate ok. Still the DB could have additional
wheels and an addition would break the constraint. If document does not have 4 wheels it would fail unless
also getting information from the DB.

What if the document contains import statements. Those objects would be looked up in the DB and made available
in the evaluator. 

There would however be no need to import already existing "wheels".

The only easily implementable solution is to NOT provide cross object validation - neither for compositions or associations.
But that sucks and defeats the purpose of invariants to a great degree.

Idea... transform expression to a query.
Idea... if validation needs to query it could either fail, or command could have option to treat the DB as being empty
and thus only validate what is available in the document. (Which is useful in many cases and certainly works for only
per instance constraints.)

Even if querying, it isn't safe since validation does not take place in a transaction. Thus what is read could
have been upated or deleted by the time the operation gets to actually inserting data. Must have one process
at a time doing updating to work around this. (Note that weaviate does not have transactions).

Neo4j offers manually controlled transactions. The Tx is started, operations are performed and the transaction can be rolled back.
Question is what the size of the tx can be. Checked, does not seem to have one except running out of memory.
So, it is doable to run queries to perform validation. The task then becomes to figure out what is needed in terms
of a graph rooted at the instance in question - this can be computed from the expressions.

```
$self.HAS_Foo..HAS_Bar..HAS_Fee..foo and via intermediates like
$self.HAS_Foos-> All {$0..HAS_Bar..HAS_Fee..foo}
```
To do this type inference must find the type of $0 in the example above.
It must also construct the query to get the graph.
The graph must be organized so the validator can traverse it as it evaluates the invariants.

Outline:
* graph is validated with all basic rules and all invariants that only require each instance.
* During this process, all other invariants are collected
* Type is inferred to compute the graph to query.

For now. Do not implement `.` operator for relations, and do not implement `..`.



