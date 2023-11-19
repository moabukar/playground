# First stage: Build the image with 'sh' installed
FROM amazon/cloudwatch-agent as builder

# Install 'sh' package
RUN apk update && apk add --no-cache gcc musl-dev
RUN apk add --no-cache --virtual .build-deps git
RUN git clone https://github.com/rofl0r/musl-fts.git /musl-fts
RUN cd /musl-fts && make && cp /musl-fts/musl-fts /bin/sh

# Second stage: Use BusyBox as the base image
FROM scratch

# Copy 'sh' from the first stage to the second stage
COPY --from=builder /bin/sh /bin/sh

# Set the entry point or default command for your container
CMD ["/bin/sh"]
