### Package: domain_util

This package provides utility functions for domain name manipulation and parsing.

#### Functions:

1. **SplitFldDomainAndSubDomain(domain string) (fldDomain string, subDomainName string, err error)**
    - **Description:** Splits a domain name into its first-level domain (FLD) and subdomain parts.
    - **Parameters:**
        - `domain (string)`: The domain name to be split.
    - **Returns:**
        - `fldDomain (string)`: The first-level domain part of the domain.
        - `subDomainName (string)`: The subdomain part of the domain.
        - `err (error)`: Any error that occurs during the process.
    - **Example:**
      ```go
      fldDomain, subDomain, err := domain_util.SplitFldDomainAndSubDomain("sub.example.com")
      if err != nil {
          // handle error
      }
      fmt.Println(fldDomain)  // Output: "example.com"
      fmt.Println(subDomain) // Output: "sub"
      ```

2. **FldDomain(domain string) (string, error)**
    - **Description:** Extracts the first-level domain part from a domain name.
    - **Parameters:**
        - `domain (string)`: The domain name to extract the FLD from.
    - **Returns:**
        - `string`: The first-level domain part of the domain.
        - `error`: Any error that occurs during the process.
    - **Example:**
      ```go
      fldDomain, err := domain_util.FldDomain("sub.example.com")
      if err != nil {
          // handle error
      }
      fmt.Println(fldDomain) // Output: "example.com"
      ```

3. **FldDomainIgnoreError(domain string) string**
    - **Description:** Extracts the first-level domain part from a domain name, ignoring any errors.
    - **Parameters:**
        - `domain (string)`: The domain name to extract the FLD from.
    - **Returns:**
        - `string`: The first-level domain part of the domain or an empty string if an error occurs.
    - **Example:**
      ```go
      fldDomain := domain_util.FldDomainIgnoreError("sub.example.com")
      fmt.Println(fldDomain) // Output: "example.com"
      ```

4. **IsFldDomain(domain string) (bool, error)**
    - **Description:** Determines if the given domain is a first-level domain.
    - **Parameters:**
        - `domain (string)`: The domain to check.
    - **Returns:**
        - `bool`: True if the domain is a FLD, false otherwise.
        - `error`: Any error that occurs during the process.
    - **Example:**
      ```go
      isFld, err := domain_util.IsFldDomain("example.com")
      if err != nil {
          // handle error
      }
      fmt.Println(isFld) // Output: true
      ```

5. **NewHostEntry(host string) (*HostEntry, error)**
    - **Description:** Parses a domain name and separates its first-level domain and subdomain information into a `HostEntry` struct.
    - **Parameters:**
        - `host (string)`: The domain name to parse.
    - **Returns:**
        - `*HostEntry`: A pointer to a `HostEntry` struct containing the parsed information.
        - `error`: Any error that occurs during the process.
    - **Example:**
      ```go
      entry, err := domain_util.NewHostEntry("sub.example.com")
      if err != nil {
          // handle error
      }
      fmt.Println(entry.Domain) // Output: "example.com"
      ```

#### Types:

1. **HostEntry**
    - **Description:** A struct representing a host entry with its domain information.
    - **Fields:**
        - `Host (string)`: The original host string.
        - `Tld (string)`: The top-level domain.
        - `Domain (string)`: The first-level domain.
        - `SubName (string)`: The subdomain.
        - `WildcardBase (string)`: The base for wildcard records.
        - `WildcardRecords []Record`: A slice of wildcard records.
        - `Wildcard (bool)`: Indicates if the host has wildcard records.

2. **Record**
    - **Description:** A struct representing a DNS record.
    - **Fields:**
        - `Type (string)`: The type of the DNS record.
        - `Value (string)`: The value of the DNS record.

#### Notes:
- The `isASCII` function is defined twice with the same implementation. It should be removed or refactored to avoid redundancy.
- The `HostEntry` struct contains fields that are not used in the provided functions (`WildcardRecords` and `Wildcard`). These fields should be either used or removed to maintain clarity and avoid confusion.

This documentation provides a clear understanding of each function's purpose, parameters, return values, and usage examples, as well as descriptions of the defined types.
