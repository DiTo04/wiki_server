FROM scratch
MAINTAINER David Tennander <davidten@kth.se>
ADD wiki_server /wiki_server
ADD static /static
ADD templates /templates
ADD pages /pages
ENTRYPOINT ["/wiki_server"]