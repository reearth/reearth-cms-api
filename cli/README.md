# Re:Earth CMS CLI

A command-line interface for interacting with Re:Earth CMS API.

## Installation

```bash
go install github.com/reearth/reearth-cms-api/cli/cmd@latest
```

Or build from source:

```bash
cd cli
go build -o cms ./cmd/
```

## Configuration

### Environment Variables

| Variable | Description |
|----------|-------------|
| `REEARTH_CMS_BASE_URL` | CMS API base URL |
| `REEARTH_CMS_TOKEN` | API token |
| `REEARTH_CMS_WORKSPACE` | Workspace ID (optional) |

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

# Create an item
cms items create -m <model-id> -f '[{"key":"title","type":"text","value":"Hello"}]'
cms items create -m <model-key> -p <project-id> -f '<json>'  # key-based access

# Update an item
cms items update <item-id> -f '[{"key":"title","type":"text","value":"Updated"}]'

# Delete an item
cms items delete <item-id>
```

### Assets

```bash
# Get an asset by ID
cms assets get <asset-id>

# Upload an asset (signed URL, recommended for large files)
cms assets upload -p <project-id> -f /path/to/file

# Upload an asset (direct upload)
cms assets upload -p <project-id> -f /path/to/file --direct

# Upload an asset from URL
cms assets upload-url -p <project-id> -u https://example.com/image.png

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
# Set environment variables
export REEARTH_CMS_BASE_URL=https://api.cms.example.com
export REEARTH_CMS_TOKEN=your-api-token

# List all models
cms models list -p my-project

# Get items with JSON output
cms items list -m my-model --json id,fields

# Upload a file
cms assets upload -p my-project -f ./image.png

# Create an item with fields
cms items create -m my-model -f '[
  {"key": "title", "type": "text", "value": "My Title"},
  {"key": "description", "type": "textarea", "value": "Description here"}
]'
```

## License

Apache-2.0
