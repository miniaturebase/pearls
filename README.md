# pearls

A set of reusable terminal UI components for the Bubbletea framework.

🚧 _WIP!_ 🚧

## `types`

Essentially these are the models that contain backing data and supporting methods for a given component.

## `pkg`

Each component has it's own package where functions are exposed to create new instances of a component in various states and configurations. These functions will return the appropriate struct type defined in the the `types` directory. For example, the `checkbox.New()` function will return a `pearls.Checkbox`!
