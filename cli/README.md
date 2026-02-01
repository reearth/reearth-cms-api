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
| `REEARTH_CMS_SAFE_MODE` | Disable destructive operations (update/delete) | `false` |

### Command-line Flags

All commands support global flags:

```
--base-url string    CMS API base URL
--token string       API token
-w, --workspace string   Workspace ID
--json string        Output as JSON (optionally specify fields: --json id,name)
```

## Commands

### Models

```bash
# List models in a project
cms models list -p <project-id-or-alias>
cms models list -p <project-id-or-alias> --page 1 --per-page 20

# Get a model by ID
cms models get <model-id>

# Get a model by key (requires project)
cms models get <model-key> -p <project-id-or-alias>
```

### Items

```bash
# List items in a model
cms items list -m <model-id>
cms items list -m <model-key> -p <project-id>  # key-based access
cms items list -m <model-id> --page 1 --per-page 20 --asset

# Get an item by ID
cms items get <item-id>
cms items get <item-id> --asset

# Create an item (use -k/-t/-v for each field)
cms items create -m <model-id> -k title -t text -v "Hello"
cms items create -m <model-id> -k title -t text -v "Hello" -k count -t number -v 10
cms items create -m <model-key> -p <project-id> -k title -t text -v "Hello"  # key-based access

# Create an item with metadata fields (use -K/-T/-V for metadata)
cms items create -m <model-id> -k title -t text -v "Hello" -K status -T select -V "published"

# Update an item (requires confirmation)
cms items update <item-id> -k title -t text -v "Updated"
cms items update <item-id> -k title -t text -v "Updated" -y  # skip confirmation

# Update an item with metadata
cms items update <item-id> -K status -T select -V "draft"

# Delete an item (requires confirmation)
cms items delete <item-id>
cms items delete <item-id> -y  # skip confirmation
```

### Assets

```bash
# Get an asset by ID
cms assets get <asset-id>

# Create an asset from file (signed URL, recommended for large files)
cms assets create -p <project-id> -f /path/to/file

# Create an asset from file (direct upload)
cms assets create -p <project-id> -f /path/to/file --direct

# Create an asset from URL
cms assets create -p <project-id> -u https://example.com/image.png

# Output asset content to stdout
cms assets cat <asset-id>

# Copy asset content to a file
cms assets cp <asset-id> /path/to/destination
```

### Comments

```bash
# Add a comment to an item
cms comments item <item-id> -c "Comment content"

# Add a comment to an asset
cms comments asset <asset-id> -c "Comment content"
```

## Output Formats

### Table (default)

```bash
cms models list -p my-project
```

### JSON

```bash
# Full JSON output
cms models list -p my-project --json

# Select specific fields
cms models list -p my-project --json id,name,key
```

## Examples

```bash
# Set environment variables (base URL defaults to https://api.cms.reearth.io)
export REEARTH_CMS_TOKEN=your-api-token

# List all models
cms models list -p my-project

# Get items with JSON output
cms items list -m my-model --json id,fields

# Create an asset from file
cms assets create -p my-project -f ./image.png

# Create an item with multiple fields
cms items create -m my-model \
  -k title -t text -v "My Title" \
  -k description -t textarea -v "Description here"
```

## License

Apache-2.0
