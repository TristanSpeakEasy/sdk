# CreateProjectGitComments

## Example Usage

```typescript
import { CreateProjectGitComments } from "@vercel/sdk/models/createprojectop.js";

let value: CreateProjectGitComments = {
  onPullRequest: true,
  onCommit: true,
};
```

## Fields

| Field                                            | Type                                             | Required                                         | Description                                      |
| ------------------------------------------------ | ------------------------------------------------ | ------------------------------------------------ | ------------------------------------------------ |
| `onPullRequest`                                  | *boolean*                                        | :heavy_check_mark:                               | Whether the Vercel bot should comment on PRs     |
| `onCommit`                                       | *boolean*                                        | :heavy_check_mark:                               | Whether the Vercel bot should comment on commits |