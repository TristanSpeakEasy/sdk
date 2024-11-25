# UpdateProjectTrustedIps1

## Example Usage

```typescript
import { UpdateProjectTrustedIps1 } from "@vercel/sdk/models/operations/updateproject.js";

let value: UpdateProjectTrustedIps1 = {
  deploymentType: "prod_deployment_urls_and_all_previews",
  addresses: [
    {
      value: "<value>",
    },
  ],
  protectionMode: "additional",
};
```

## Fields

| Field                                                                                                                | Type                                                                                                                 | Required                                                                                                             | Description                                                                                                          |
| -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- |
| `deploymentType`                                                                                                     | [operations.UpdateProjectTrustedIpsDeploymentType](../../models/operations/updateprojecttrustedipsdeploymenttype.md) | :heavy_check_mark:                                                                                                   | N/A                                                                                                                  |
| `addresses`                                                                                                          | [operations.UpdateProjectTrustedIpsAddresses](../../models/operations/updateprojecttrustedipsaddresses.md)[]         | :heavy_check_mark:                                                                                                   | N/A                                                                                                                  |
| `protectionMode`                                                                                                     | [operations.UpdateProjectTrustedIpsProtectionMode](../../models/operations/updateprojecttrustedipsprotectionmode.md) | :heavy_check_mark:                                                                                                   | N/A                                                                                                                  |