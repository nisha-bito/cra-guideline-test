# Error handling rules
- Wrap every returned error with context using fmt.Errorf("...: %w", err).
- Never compare errors with ==; use errors.Is or errors.As.
