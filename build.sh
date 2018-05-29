set -x

TAG=$1
if [[ -z "${TAG}" ]]
then
	TAG="v$(date '+%Y-%m-%d')"
fi

PREFIX=gcr.io/hots-cockroach/website
IMG=${PREFIX}:${TAG}
LATEST=${PREFIX}:latest

echo building $TAG
echo $IMG

docker pull $IMG 2>/dev/null >/dev/null
if [[ $? -eq 0 ]]
then
	echo 'image already exists'
	exit 1
fi

set -e

docker build -t $IMG .
docker tag $IMG $LATEST
docker push $IMG
docker push $LATEST
kubectl set image deployment/website "website=${IMG}"
kubectl get po | grep updatedb | awk '{print $1}' | xargs kubectl delete po
