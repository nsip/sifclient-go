FINAL:
    * New repository
    * New README

https://irahardianto.github.io/service-pattern-go/

see also resty

Learn how

Compare Go Request instead of Resty

Extract the returned TOKEN from the Environment create
Work out how to add that back in as Basic header
OR actually get session token.

LEARNING:
    * Videos

TESTING:
    * Encoding Environment XML
    * Reading Environment data (note for future)
Important Stuff to complete for initial release FEB 2020

    * Add Root CA X509 to allow for HTTPS
        - Worked around
    * Proper documentation GoDoc
    * Proper examples
    * Move test/* to proper X_test.go Format
    * Cache client
    * Remove hard coded URLs, use Environment returned
    * Paginated results
    * Example with save to file
    * How to deal with multiple pages into single output/File
    * How to combine the XML returned (removing the srapper header/footer ?)


Future versions: - post June 2020

    * Examples with streams (must still work with pagination)
    * Automatic login (Create environment) if it is necessary
        (allowing you just to got to Get, and let the library cope
        with timeouts etc)



older - to review

    file... write( client.Get("StudentPersonals") )

        reader, err := sifclient.GetAll("StudentPersonals")

            ... pagination

            <SPs>
                <sp>
                </sp>
                <sp>
                </sp>
            </SPs>
            <SPs>
                <sp>
                </sp>
                <sp>
                </sp>
            </SPs>
            <SPs>
                <sp>
                </sp>
                <sp>
                </sp>
            </SPs>


        Next....

            pager := sifclient.GetStart(...)
            str, err := pager.GetPage()

Multiple Page issue

    1. Store to file, and fix the multiple wrappers
        . Reader utility
    2. Paginated next style (similar to DB query)
        .
    3. Stream reader
        .
