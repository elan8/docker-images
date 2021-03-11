docker run --rm --volume "$(pwd):/workspace/project" -w /workspace jevon82/buf-extended:latest \
/bin/sh -c "cp ./project/buf.yaml ./googleapis/ && cd googleapis  && buf ls-files"