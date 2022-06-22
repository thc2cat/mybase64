# mybase64

mybase64 is a mass base64 decoder that read stdin.

- If input miss a "=" at the end, it will try to add and decode it.
- If output contains non ASCII char, it will produce a warning instead
