# UpdateCheckRequest

## Example Usage

```typescript
import { UpdateCheckRequest } from "@vercel/sdk/models/updatecheckop.js";

let value: UpdateCheckRequest = {
  deploymentId: "dpl_2qn7PZrx89yxY34vEZPD31Y9XVj6",
  checkId: "check_2qn7PZrx89yxY34vEZPD31Y9XVj6",
  teamId: "team_1a2b3c4d5e6f7g8h9i0j1k2l",
  slug: "my-team-url-slug",
};
```

## Fields

| Field                                                                | Type                                                                 | Required                                                             | Description                                                          | Example                                                              |
| -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- |
| `deploymentId`                                                       | *string*                                                             | :heavy_check_mark:                                                   | The deployment to update the check for.                              | dpl_2qn7PZrx89yxY34vEZPD31Y9XVj6                                     |
| `checkId`                                                            | *string*                                                             | :heavy_check_mark:                                                   | The check being updated                                              | check_2qn7PZrx89yxY34vEZPD31Y9XVj6                                   |
| `teamId`                                                             | *string*                                                             | :heavy_minus_sign:                                                   | The Team identifier to perform the request on behalf of.             | team_1a2b3c4d5e6f7g8h9i0j1k2l                                        |
| `slug`                                                               | *string*                                                             | :heavy_minus_sign:                                                   | The Team slug to perform the request on behalf of.                   | my-team-url-slug                                                     |
| `requestBody`                                                        | [models.UpdateCheckRequestBody](../models/updatecheckrequestbody.md) | :heavy_check_mark:                                                   | N/A                                                                  |                                                                      |