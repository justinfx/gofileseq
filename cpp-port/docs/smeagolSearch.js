/*
This is a copy of a file from the docs.wetafx.co.nz server.
Please keep the two files in synx if possible.
*/

function searchAll() {
    var baseUrl = "https://spongebob.wetafx.co.nz/search?proxystylesheet=default_frontend&q=";
    var searchTerms = document.getElementById("searchTerms").value;
    console.log("Smeagol will search for... "+baseUrl+searchTerms);
    window.location=baseUrl+searchTerms;
}

/* This is a simple localised search for Smeagol - would require extending for non-standard or complex urls */
/* next steps: pass url fragments to check as a parameter; then check for any of a list of fragments. */
/* Rules are cumulative (that is, an url must be _in_ the whitelist and _not in_ the black list to Smeagol. */
function showHideByUrl(whitelisturls=[],blacklisturls=[]) {
    console.log("Started shbu");
    var encodedSearch = encodeURI(document.domain)+encodeURI(window.location.pathname);
    var doxygenSearch = document.querySelector("#MSearchBox");
    var smeagolSearch = document.getElementsByClassName("headerSearch")[0];
    var showSmeagol = false;
    smeagolSearch.style.display = "none";
    doxygenSearch.style.display = "none";
    console.log("blacklisturls: "+blacklisturls+"; whitelisturls: "+whitelisturls);
    console.log("encodedURI: "+encodedSearch);
    for ( i = 0; i < whitelisturls.length; i++ ) {
        /* if we find the url in the whitelist, set smeagol search to true */
        if (encodedSearch.search(whitelisturls[i]) > -1) {
            showSmeagol = true;
        }
/*        console.log("iterating white: "+whitelisturls[i]+" showSmeagol = "+showSmeagol);*/
    }
    for ( i = 0; i < blacklisturls.length; i++ ) {
        if (encodedSearch.search(blacklisturls[i]) > -1 ) {
            console.log("blacklisted");
            showSmeagol = false;
        }
/*        console.log("iterating black: "+blacklisturls[i]+" showSmeagol = "+showSmeagol);*/
    }
    if ( showSmeagol == true ) {
        console.log("smeagolSearch!");
        smeagolSearch.style.display = "inline";
    } else {
        console.log("doxygenSearch!");
        doxygenSearch.style.display = "inline";
    }
    console.log("Finished shbu");
}

/* This is a simple localised search for Smeagol - would require extending for non-standard or complex urls */
function searchLocalDocs() {
    /* where to find smeagol, and required guff for search url. */
    var baseUrl = "https://spongebob.wetafx.co.nz/search?&proxystylesheet=default_frontend&filter=0";
    /* limit search to current url and below. */
/*    var searchLimit = "&q=inurl%3A";  using site: instead as it doesn't require backslashing fullstops */
    var searchLimit = "&q=site%3A";
    /* note: encodeURI doesn't encode things such as: ./!*:? so this is just-in-case. */
    var encodedSearch = encodeURI(window.location.pathname);
    var searchPathArray = encodedSearch.split( '/' );
    searchPathArray.pop();
    var searchPath = searchPathArray.join("%2F");
    /* If you want to use inurl:, you need to do:
       var tmpVariable = searchPath.replace(/\./g, '%5C%2E');
       It's required, as Smeagol won't honor fullstops in urls for inurl: (other than host name). 
    /* This would be where we'd encode the other reserved characters, esp ?+&. */
    var searchDomain = window.location.host+searchPath;
    searchLimit += searchDomain;
    /* collect terms to search for */
    var searchTerms = document.getElementById("searchTerms").value;
    console.log("Asking Smeagol to search for... "+baseUrl+searchLimit+"+"+searchTerms);
    /* go to smeagol with search limit and terms */
    window.location=baseUrl+searchLimit+"+"+searchTerms;
}
