## grafanactl dev lint rules

List available linter rules

### Synopsis

List available linter rules.

```
grafanactl dev lint rules [flags]
```

### Examples

```

	# List built-in rules:

	grafanactl dev lint rules

	# List built-in and custom rules:

	grafanactl dev lint rules -r ./custom-rules

```

### Options

```
  -h, --help                help for rules
      --json string         Comma-separated list of fields to include in JSON output, or '?' to discover available fields
  -o, --output string       Output format. One of: json, yaml (default "yaml")
  -r, --rules stringArray   Path to custom rules.
```

### Options inherited from parent commands

```
      --agent            Enable agent mode (JSON output, no color). Auto-detected from CLAUDECODE, CLAUDE_CODE, CURSOR_AGENT, GITHUB_COPILOT, AMAZON_Q, or GRAFANACTL_AGENT_MODE env vars.
      --config string    Path to the configuration file to use
      --context string   Name of the context to use
      --no-color         Disable color output
      --no-truncate      Disable table column truncation (auto-enabled when stdout is piped)
  -v, --verbose count    Verbose mode. Multiple -v options increase the verbosity (maximum: 3).
```

### SEE ALSO

* [grafanactl dev lint](grafanactl_dev_lint.md)	 - Lint Grafana resources

