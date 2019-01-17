# Wally - Website Walker

Wally walks websites verifying the health and performance of the
website as a whole, all individual (accessible) pages, like things
that including:

- site/page availability
- site/page performance 
- site/page link checking
- site/page structure

An Example

```bash

% wally -depth 1 http://sierrahydrographics.com
Pages ...
	http://gardenpassages.com/custom-wood-gates
	http://gardenpassages.com/wood-gate-maintenance
	http://gardenpassages.com/about-garden-passages
	http://gardenpassages.com/old-gate-gallery
	http://gardenpassages.com/finishing-your-gates
```


## Wally 

Wally maybe asked to walk a site asynchronously, in that case the
request is associated with a Job token, that can be used to locate the
results once they have completed.

