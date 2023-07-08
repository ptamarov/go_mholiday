

## Takeaways 

1. The best way to deal with errors is to make them impossible.
2. Try to reduce edge cases that are hard to test or debug.
3. Design abstractions so that operations are safe.

## Proactively prevent problems

1. Every piece of data in your software should start its life in a valid state.
2. Every transformation should leave it in a valid state.

Guidelines:

- Break large programs into small pieces you can understand
- Hide information to reduce corruption
- Avoid clever code and side effects 
- Avoid _unsafe_ operations
- Assert your invariants
- _Never_ ignore errors
- _Never_ accept input from a user without validation
- Test! You can't test enough. 