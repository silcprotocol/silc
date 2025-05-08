import pytest

from .network import setup_silc, setup_silc_rocksdb, setup_geth


@pytest.fixture(scope="session")
def silc(tmp_path_factory):
    path = tmp_path_factory.mktemp("silc")
    yield from setup_silc(path, 26650)


@pytest.fixture(scope="session")
def silc_rocksdb(tmp_path_factory):
    path = tmp_path_factory.mktemp("silc-rocksdb")
    yield from setup_silc_rocksdb(path, 20650)


@pytest.fixture(scope="session")
def geth(tmp_path_factory):
    path = tmp_path_factory.mktemp("geth")
    yield from setup_geth(path, 8545)


@pytest.fixture(scope="session", params=["silc", "silc-ws"])
def silc_rpc_ws(request, silc):
    """
    run on both silc and silc websocket
    """
    provider = request.param
    if provider == "silc":
        yield silc
    elif provider == "silc-ws":
        silc_ws = silc.copy()
        silc_ws.use_websocket()
        yield silc_ws
    else:
        raise NotImplementedError


@pytest.fixture(scope="module", params=["silc", "silc-ws", "silc-rocksdb", "geth"])
def cluster(request, silc, silc_rocksdb, geth):
    """
    run on silc, silc websocket,
    silc built with rocksdb (memIAVL + versionDB)
    and geth
    """
    provider = request.param
    if provider == "silc":
        yield silc
    elif provider == "silc-ws":
        silc_ws = silc.copy()
        silc_ws.use_websocket()
        yield silc_ws
    elif provider == "geth":
        yield geth
    elif provider == "silc-rocksdb":
        yield silc_rocksdb
    else:
        raise NotImplementedError


@pytest.fixture(scope="module", params=["silc", "silc-rocksdb"])
def silc_cluster(request, silc, silc_rocksdb):
    """
    run on silc default build &
    silc with rocksdb build and memIAVL + versionDB
    """
    provider = request.param
    if provider == "silc":
        yield silc
    elif provider == "silc-rocksdb":
        yield silc_rocksdb
    else:
        raise NotImplementedError
