# ProofChain-Store

Backend service for secure storage and retrieval of already-encrypted documents for the **[ProofChain](https://www.github.com/Raaffs/ProofChain)** platform.  
Acts as a secure interface to MongoDB, keeping database credentials off users’ desktops.

## Features

- Store and retrieve documents that are **already encrypted**—ProofChain-Store does **not perform encryption** itself.
- Minimal admin-only routes for metadata management:
  - Add registered institutions to the database.
  - Map document types to verifying institutions.

## Overview 

ProofChain-Store is designed to separate storage concerns from the desktop application:

- **Encrypted Documents:** The backend only stores documents that have been encrypted on the client side.  
- **Admin Metadata:** Only a few routes are exposed for contract owners to maintain institution and document type mappings.  

