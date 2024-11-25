# GitSource

Defines the Git Repository source to be deployed. This property can not be used in combination with `files`.


## Supported Types

### `operations.One`

```typescript
const value: operations.One = {
  ref: "<value>",
  repoId: 1023.17,
  type: "github",
};
```

### `operations.Two`

```typescript
const value: operations.Two = {
  org: "<value>",
  ref: "<value>",
  repo: "<value>",
  type: "github",
};
```

### `operations.Three`

```typescript
const value: operations.Three = {
  projectId: 9829.90,
  ref: "<value>",
  type: "gitlab",
};
```

### `operations.Four`

```typescript
const value: operations.Four = {
  ref: "<value>",
  repoUuid: "<id>",
  type: "bitbucket",
};
```

### `operations.Five`

```typescript
const value: operations.Five = {
  owner: "<value>",
  ref: "<value>",
  slug: "<value>",
  type: "bitbucket",
};
```

