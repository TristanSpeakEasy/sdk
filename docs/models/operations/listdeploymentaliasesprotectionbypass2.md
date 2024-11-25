# ListDeploymentAliasesProtectionBypass2

The protection bypass for the alias

## Example Usage

```typescript
import { ListDeploymentAliasesProtectionBypass2 } from "@vercel/sdk/models/operations/listdeploymentaliases.js";

let value: ListDeploymentAliasesProtectionBypass2 = {
  createdAt: 6935.92,
  lastUpdatedAt: 6734.58,
  lastUpdatedBy: "<value>",
  access: "granted",
  scope: "user",
};
```

## Fields

| Field                                                                                                                                        | Type                                                                                                                                         | Required                                                                                                                                     | Description                                                                                                                                  |
| -------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------- |
| `createdAt`                                                                                                                                  | *number*                                                                                                                                     | :heavy_check_mark:                                                                                                                           | N/A                                                                                                                                          |
| `lastUpdatedAt`                                                                                                                              | *number*                                                                                                                                     | :heavy_check_mark:                                                                                                                           | N/A                                                                                                                                          |
| `lastUpdatedBy`                                                                                                                              | *string*                                                                                                                                     | :heavy_check_mark:                                                                                                                           | N/A                                                                                                                                          |
| `access`                                                                                                                                     | [operations.ListDeploymentAliasesProtectionBypassAccess](../../models/operations/listdeploymentaliasesprotectionbypassaccess.md)             | :heavy_check_mark:                                                                                                                           | N/A                                                                                                                                          |
| `scope`                                                                                                                                      | [operations.ListDeploymentAliasesProtectionBypassAliasesScope](../../models/operations/listdeploymentaliasesprotectionbypassaliasesscope.md) | :heavy_check_mark:                                                                                                                           | N/A                                                                                                                                          |