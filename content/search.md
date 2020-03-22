---
title: "Search Results"
sitemap:
  priority : 0.1
---

{{< safe-html >}}
<div id="search-results"></div>
<script id="search-result-template" type="text/x-js-template">
    <div id="summary-${key}">
        <h3><a href="${link}">${title}</a></h3>
        <p>${snippet}</p>
    </div>
</script>
{{< /safe-html >}}