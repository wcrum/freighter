# Freighter

## Airgap Swiss Army Knife

`Freighter` is a **community fork of [Rancher Government Hauler](https://hauler.dev)**, created and maintained by [William Crum](https://github.com/wcrum).  

This project continues the mission of **simplifying the airgap experience** without requiring operators to adopt a specific workflow.  
Like its upstream, `Hauler` represents assets (images, charts, files, etc.) as content and collections, allowing operators to easily fetch, store, package, and distribute them with declarative manifests or through the command line.

### Disclaimer
This is an **independent fork** of the original [Hauler](https://hauler.dev) project.  
It is **not affiliated with, endorsed by, or sponsored by Rancher Government Solutions or the original Hauler maintainers**.  

---

## How Freighter Works

`Freighter` builds on Hauler's foundation by:

- Storing contents and collections as **OCI Artifacts**.  
- Serving contents and collections with an **embedded registry and fileserver**.  
- Supporting inspection and storage of various **non-image OCI Artifacts**.  

---

## Recent Changes

This fork starts from **Hauler v1.2.0** and introduces ongoing modifications by William Crum.  
Please refer to [CHANGELOG.md](./CHANGELOG.md) for specific differences between `Hauler` and upstream `Freighter`.

### Inherited from Hauler v1.2.0:

- **API version upgrade** from `v1alpha1` to `v1`.  
  - `v1alpha1` is now deprecated and will be removed in a future release.  
  - Logging notices are displayed when using `v1alpha1`, e.g.:  
    ```
    !!! DEPRECATION WARNING !!! apiVersion [v1alpha1] will be removed in a future release !!! DEPRECATION WARNING !!!
    ```
---
- Updated behavior of `store load`:  
  - Defaults to loading a `haul` with the name `haul.tar.zst`.  
  - Requires `--filename/-f` to load a haul with a different name.  
  - Supports multiple hauls with multiple flags.  
  - Example:  
    ```bash
    freighter store load --filename hauling-hauls.tar.zst
    ```
---
- Updated behavior of `store sync`:  
  - Defaults to syncing a `manifest` named `freighter-manifest.yaml`.  
  - Requires `--filename/-f` to sync other manifest names.  
  - Supports multiple manifests with multiple flags.  
  - Example:  
    ```bash
    freighter store sync --filename hauling-hauls-manifest.yaml
    ```
---
For known limits, issues, and notices, see the [upstream Hauler documentation](https://docs.freighter.dev/docs/known-limits).  
Freighter-specific changes will be tracked separately.

---

## Installation

## License & Notices
- This project includes code from Rancher Government Hauler, licensed under the Apache License 2.0.
- See NOTICE for required attributions from upstream.

