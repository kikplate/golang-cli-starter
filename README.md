# golang-cli-starter Kikplate Template

This repository is a Kikplate template source for a production-ready Go CLI starter.

## Generate from local template

```bash
kik generate --template . --output-dir ./generated
```

## Generate with custom values

```bash
kik generate --template . --set projectName=my-cli --set modulePath=github.com/acme/my-cli --set binaryName=mycli --output-dir ./generated
```

## Manifest

- Plate manifest: `plate.yaml`
- Template files: `templates/`

## Notes

- `plate.yaml` uses raw GitHub template URLs for all generated files.
- The generated project can be customized through schema values such as `modulePath`, `binaryName`, and `envPrefix`.
