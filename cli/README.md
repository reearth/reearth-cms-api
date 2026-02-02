# Re:Earth CMS CLI

A command-line interface for interacting with Re:Earth CMS API.

## Installation

```bash
go install github.com/reearth/reearth-cms-api/cli/reearth-cms@latest
```

Or build from source:

```bash
cd cli
go build -o reearth-cms ./reearth-cms/
```

## Configuration

### Environment Variables

Environment variables can also be set in a `.env` file in the current directory.

| Variable | Description | Default |
|----------|-------------|---------|
| `REEARTH_CMS_BASE_URL` | CMS API base URL | `https://api.cms.reearth.io` |
| `REEARTH_CMS_TOKEN` | API token | (required) |
| `REEARTH_CMS_WORKSPACE` | Workspace ID | (optional) |
| `REEARTH_CMS_PROJECT` | Project ID or alias | (optional) |
| `REEARTH_CMS_SAFE_MODE` | Disable destructive operations (update/delete) | `false` |

### Command-line Flags

All commands support global flags:

```
--base-url string        CMS API base URL
--token string           API token
-w, --workspace string   Workspace ID
-p, --project string     Project ID or alias
--json string            Output as JSON (optionally specify fields: --json id,name)
```

## Commands

### Model

```bash
# List models in a project
reearth-cms model list -p <project-id-or-alias>
reearth-cms model list -p <project-id-or-alias> --page 1 --per-page 20

# Get a model by ID
reearth-cms model get <model-id>

# Get a model by key (requires project)
reearth-cms model get <model-key> -p <project-id-or-alias>
```

### Item

```bash
# List items in a model
reearth-cms item list -m <model-id>
reearth-cms item list -m <model-key> -p <project-id>  # key-based access
reearth-cms item list -m <model-id> --page 1 --per-page 20 --asset

# Get an item by ID (asset data is always included)
reearth-cms item get <item-id>

# Create an item (use -k/-t/-v for each field)
reearth-cms item create -m <model-id> -k title -t text -v "Hello"
reearth-cms item create -m <model-id> -k title -t text -v "Hello" -k count -t number -v 10
reearth-cms item create -m <model-key> -p <project-id> -k title -t text -v "Hello"  # key-based access

# Create an item with metadata fields (use -K/-T/-V for metadata)
reearth-cms item create -m <model-id> -k title -t text -v "Hello" -K status -T select -V "published"

# Update an item (requires confirmation)
reearth-cms item update <item-id> -k title -t text -v "Updated"
reearth-cms item update <item-id> -k title -t text -v "Updated" -y  # skip confirmation

# Update an item with metadata
reearth-cms item update <item-id> -K status -T select -V "draft"

# Delete an item (requires confirmation)
reearth-cms item delete <item-id>
reearth-cms item delete <item-id> -y  # skip confirmation
```

### Asset

```bash
# Get an asset by ID
reearth-cms asset get <asset-id>

# Create an asset from file (signed URL, recommended for large files)
reearth-cms asset create -p <project-id> -f /path/to/file

# Create an asset from file (direct upload)
reearth-cms asset create -p <project-id> -f /path/to/file --direct

# Create an asset from URL
reearth-cms asset create -p <project-id> -u https://example.com/image.png

# Output asset content to stdout
reearth-cms asset cat <asset-id>

# Copy asset content to a file
reearth-cms asset cp <asset-id> /path/to/destination
```

### Comment

```bash
# Add a comment to an item
reearth-cms comment item <item-id> -c "Comment content"

# Add a comment to an asset
reearth-cms comment asset <asset-id> -c "Comment content"
```

## Output Formats

### YAML (default)

```bash
reearth-cms model list -p my-project
```

### JSON

```bash
# Full JSON output
reearth-cms model list -p my-project --json=1

# Select specific fields
reearth-cms model list -p my-project --json=id,name,key
```

## Examples

### Using .env file

Create a `.env` file in your working directory:

```bash
REEARTH_CMS_BASE_URL=https://api.cms.reearth.io
REEARTH_CMS_TOKEN=your-api-token
REEARTH_CMS_PROJECT=my-project
REEARTH_CMS_SAFE_MODE=true
```

Then run commands without specifying project each time:

```bash
# List all models (uses REEARTH_CMS_PROJECT from .env)
reearth-cms model list

# Get items with JSON output
reearth-cms item list -m my-model --json=id,fields

# Create an asset from file
reearth-cms asset create -f ./image.png

# Create an item with multiple fields
reearth-cms item create -m my-model \
  -k title -t text -v "My Title" \
  -k description -t textarea -v "Description here"
```

### Using environment variables

```bash
export REEARTH_CMS_TOKEN=your-api-token
export REEARTH_CMS_PROJECT=my-project

reearth-cms model list
```

### Using command-line flags

```bash
reearth-cms model list -p my-project --token your-api-token
```

## License

Apache-2.0
