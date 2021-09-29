### Example

Archive, then export both development and app-store IPAs:

```yaml
steps:
- xcode-archive:
    title: Archive and export development IPA
    inputs:
    - export_method: development
- export-xcarchive:
    title: Export app-store IPA
    inputs:
    - archive_path: $BITRISE_XCARCHIVE_PATH # this env var is an output of the previous step
    - export_method: app-store
- deploy-to-bitrise-io:
```