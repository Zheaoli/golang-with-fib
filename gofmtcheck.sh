if [[ $(gofmt -d ./) ]]; then
  echo "run gofmt -w ./ to fix code"
  exit 1
else
  exit 0
fi

