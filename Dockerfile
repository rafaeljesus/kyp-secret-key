FROM scratch
MAINTAINER Rafael Jesus <rafaelljesus86@gmail.com>
ADD kyp-secret-key /kyp-secret-key
ENV KYP_SECRET_KEY_PORT="3000"
ENTRYPOINT ["/kyp-secret-key"]
