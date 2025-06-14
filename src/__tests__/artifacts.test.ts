/*
 * Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.
 */

import { expect, test } from "vitest";
import { Vercel } from "../index.js";
import { filesToByteArray, streamToByteArray } from "./files.js";
import { createTestHTTPClient } from "./testclient.js";

test("Artifacts Record Events", async () => {
  const testHttpClient = createTestHTTPClient("recordEvents");

  const vercel = new Vercel({
    serverURL: process.env["TEST_SERVER_URL"] ?? "http://localhost:18080",
    httpClient: testHttpClient,
    bearerToken: "<YOUR_BEARER_TOKEN_HERE>",
  });

  await vercel.artifacts.recordEvents({
    xArtifactClientCi: "VERCEL",
    xArtifactClientInteractive: 0,
    teamId: "team_1a2b3c4d5e6f7g8h9i0j1k2l",
    slug: "my-team-url-slug",
    requestBody: [
      {
        sessionId: "<id>",
        source: "REMOTE",
        event: "MISS",
        hash: "12HKQaOmR5t5Uy6vdcQsNIiZgHGB",
        duration: 400,
      },
    ],
  });
});

test("Artifacts Status", async () => {
  const testHttpClient = createTestHTTPClient("status");

  const vercel = new Vercel({
    serverURL: process.env["TEST_SERVER_URL"] ?? "http://localhost:18080",
    httpClient: testHttpClient,
    bearerToken: "<YOUR_BEARER_TOKEN_HERE>",
  });

  const result = await vercel.artifacts.status({
    teamId: "team_1a2b3c4d5e6f7g8h9i0j1k2l",
    slug: "my-team-url-slug",
  });
  expect(result).toBeDefined();
  expect(result).toEqual({
    status: "over_limit",
  });
});

test("Artifacts Download Artifact", async () => {
  const testHttpClient = createTestHTTPClient("downloadArtifact");

  const vercel = new Vercel({
    serverURL: process.env["TEST_SERVER_URL"] ?? "http://localhost:18080",
    httpClient: testHttpClient,
    bearerToken: "<YOUR_BEARER_TOKEN_HERE>",
  });

  const result = await vercel.artifacts.downloadArtifact({
    xArtifactClientCi: "VERCEL",
    xArtifactClientInteractive: 0,
    hash: "12HKQaOmR5t5Uy6vdcQsNIiZgHGB",
    teamId: "team_1a2b3c4d5e6f7g8h9i0j1k2l",
    slug: "my-team-url-slug",
  });
  expect(result).toBeDefined();
  expect(new Uint8Array(await streamToByteArray(result))).toEqual(
    await filesToByteArray(".speakeasy/testfiles/example.file"),
  );
});

test("Artifacts Artifact Exists", async () => {
  const testHttpClient = createTestHTTPClient("artifactExists");

  const vercel = new Vercel({
    serverURL: process.env["TEST_SERVER_URL"] ?? "http://localhost:18080",
    httpClient: testHttpClient,
    bearerToken: "<YOUR_BEARER_TOKEN_HERE>",
  });

  await vercel.artifacts.artifactExists({
    hash: "12HKQaOmR5t5Uy6vdcQsNIiZgHGB",
    teamId: "team_1a2b3c4d5e6f7g8h9i0j1k2l",
    slug: "my-team-url-slug",
  });
});

test("Artifacts Artifact Query", async () => {
  const testHttpClient = createTestHTTPClient("artifactQuery");

  const vercel = new Vercel({
    serverURL: process.env["TEST_SERVER_URL"] ?? "http://localhost:18080",
    httpClient: testHttpClient,
    bearerToken: "<YOUR_BEARER_TOKEN_HERE>",
  });

  const result = await vercel.artifacts.artifactQuery({
    teamId: "team_1a2b3c4d5e6f7g8h9i0j1k2l",
    slug: "my-team-url-slug",
    requestBody: {
      hashes: [
        "12HKQaOmR5t5Uy6vdcQsNIiZgHGB",
        "34HKQaOmR5t5Uy6vasdasdasdasd",
      ],
    },
  });
  expect(result).toBeDefined();
  expect(result).toEqual({
    "key": {
      size: 3811.11,
      taskDurationMs: 5116.13,
    },
    "key1": {
      size: 3811.11,
      taskDurationMs: 5116.13,
    },
    "key2": {
      size: 3811.11,
      taskDurationMs: 5116.13,
    },
  });
});
