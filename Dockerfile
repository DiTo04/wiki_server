FROM scratch
MAINTAINER David Tennander <davidten@kth.se>
ADD my_first_server /my_first_server
ADD static /static
ADD templates /templates
ADD pages /pages
ENTRYPOINT ["/my_first_server"]