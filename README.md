# go-container-supply-chain-example

- Signing commits with [sigstore gitsign](https://github.com/sigstore/gitsign)
- Build container with [ko](https://github.com/google/ko), SBOM as well
- Container signing with [sigstore cosign](https://github.com/sigstore/cosign)
- Fulcio Certificates - ephemeral certs w/ requester identity in certificate metadata
- Builder identification is done through the fulcio certificate
  - Certificate attests to workload identity and signs provenance
  - job_workflow_ref identifies the reusable GitHub workflow
  - repository and commit SHA identifies the calling workflow
- Rekor Verification steps
  - Download the signing certificate from the Rekor log.
  - Verify the signed provenance and signing certificate up to Fulcio: Verifies integrity, and non-falsifiability of provenance and certificate contents.
  - The builder identity from the signing certificate SAN. Verifies authenticity of the provenance and guarantees the provenance was populated.
- Attestation - https://github.com/slsa-framework/slsa/blob/main/controls/attestations.md
- Provenance - [GitHub Action Provenance Generator](https://github.com/marketplace/actions/slsa-build-provenance-action-demo)
- Scripted build - GitHub Reusable Workflow 
- CodeQL - Vulnerability Management 
- Dependabot - Vulnerability Management

# SLSA requirements

SLSA Level 3 Requirements

Source:
- Version controlled: Property of GitHub/Git
- Verified history: signed commits using gitsign
```bash
 git --no-pager log --show-signature -1
15c775104eb1eab2d305a6fdef89d46fbd7e8e99 (HEAD -> main, origin/main, origin/HEAD) tlog index: 2344075
smimesign: Signature made using certificate ID 0xee02388e292cf73b03403e41695926ecbeb1f1c2 | CN=sigstore,O=sigstore.dev
smimesign: Good signature from [strong.james.e@gmail.com]
Parsed Git signature: true
Validated Git signature: true
Located Rekor entry: true
Validated Rekor entry: true
readme up
```

Build:
- Build as code: Build using GitHub workflow in source repository
- Ephemeral environment: Property of GitHub job
- Isolated: Property of GitHub job isolation
- Parameterless: Property of GitHub workflow 
Provenance:
- Authenticated: Digital signature by ephemeral signing key provided by Fulcio
- Service generated: Builder ID attested in the certificate, trusted workflow content
- Non-falsifiable: Builder ID attested in the certificate, trusted workflow content, ephemeral key short-lived certificate


| Requirement                          | SLSA 1 | SLSA 2 | SLSA 3 | SLSA 4 |
|--------------------------------------|--------|--------|--------|--------|
| Source - [Version controlled]        |        | ✓      | ✓      | ✓      |
| Source - [Verified history]          |        |        | ✓      | ✓      |
| Source - [Retained indefinitely]     |        |        | 18 mo. | ✓      |
| Source - [Two-person reviewed]       |        |        |        | ✓      |
| Build - [Scripted build]             | ✓      | ✓      | ✓      | ✓      |
| Build - [Build service]              |        | ✓      | ✓      | ✓      |
| Build - [Build as code]              |        |        | ✓      | ✓      |
| Build - [Ephemeral environment]      |        |        | ✓      | ✓      |
| Build - [Isolated]                   |        |        | ✓      | ✓      |
| Build - [Parameterless]              |        |        |        | ✓      |
| Build - [Hermetic]                   |        |        |        | ✓      |
| Build - [Reproducible]               |        |        |        | ○      |
| Provenance - [Available]             | ✓      | ✓      | ✓      | ✓      |
| Provenance - [Authenticated]         |        | ✓      | ✓      | ✓      |
| Provenance - [Service generated]     |        | ✓      | ✓      | ✓      |
| Provenance - [Non-falsifiable]       |        |        | ✓      | ✓      |
| Provenance - [Dependencies complete] |        |        |        | ✓      |
| Common - [Security]                  |        |        |        | ✓      |
| Common - [Access]                    |        |        |        | ✓      |
| Common - [Superusers]                |        |        |        | ✓      |

  
