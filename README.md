# dispatch workflow for other's repository

Trigger github actions in multiple other repositories.

## Inputs

### `token`

**Required** github token

### `type`

**Required** event type

### `json-payload`

json format payload

## Outputs

## Example usage

```yaml:actions
  steps:
      - uses: actions/checkout@v2
      
      - name: Generate token
        id: generate_token
        uses: tibdex/github-app-token@v1
        with:
          app_id: ${{ secrets.APP_ID }}
          private_key: ${{ secrets.APP_PRIVATE_KEY }}
          
      - name: Get Json Payload
        id: get-payload
        run: |
          echo ::set-output name=PAYLOAD::$(cat test.json | jq -c .)

      - name: Dispatch other's action
        uses: goccha/dispatch-others-workflows@v0.0.1
        with:
          token: ${{ steps.generate_token.outputs.token }}
          type: deploy
          json-payload: ${{ steps.get-payload.outputs.PAYLOAD }}
```

```json:test.json
{
  "goccha/test-project1": {
    "tag": "v0.0.1"
  },
  "goccha/test-project2": {
    "tag": "v0.0.2"
  }
}
```