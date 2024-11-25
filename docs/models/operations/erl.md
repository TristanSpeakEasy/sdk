# Erl

## Example Usage

```typescript
import { Erl } from "@vercel/sdk/models/operations/updateprojectdatacache.js";

let value: Erl = {
  algo: "token_bucket",
  window: 6304.48,
  limit: 8742.88,
  keys: [
    "<value>",
  ],
};
```

## Fields

| Field                                                                                          | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `algo`                                                                                         | [operations.UpdateProjectDataCacheAlgo](../../models/operations/updateprojectdatacachealgo.md) | :heavy_check_mark:                                                                             | N/A                                                                                            |
| `window`                                                                                       | *number*                                                                                       | :heavy_check_mark:                                                                             | N/A                                                                                            |
| `limit`                                                                                        | *number*                                                                                       | :heavy_check_mark:                                                                             | N/A                                                                                            |
| `keys`                                                                                         | *string*[]                                                                                     | :heavy_check_mark:                                                                             | N/A                                                                                            |