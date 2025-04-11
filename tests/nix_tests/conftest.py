import pytest

from .network import setup_evmos, setup_evmos_rocksdb, setup_geth


@pytest.fixture(scope="session")
def silc(tmp_path_factory):
    path = tmp_path_factory.mktemp("silc")
    yield from setup_evmos(path, 26650)


@pytest.fixture(scope="session")
def evmos_rocksdb(tmp_path_factory):
    path = tmp_path_factory.mktemp("silc-rocksdb")
    yield from setup_evmos_rocksdb(path, 20650)


@pytest.fixture(scope="session")
def geth(tmp_path_factory):
    path = tmp_path_factory.mktemp("geth")
    yield from setup_geth(path, 8545)


@pytest.fixture(scope="session", params=["silc", "silc-ws"])
def evmos_rpc_ws(request, silc):
    """
    run on both silc and silc websocket
    """
    provider = request.param
    if provider == "silc":
        yield silc
    elif provider == "silc-ws":
        evmos_ws = silc.copy()
        evmos_ws.use_websocket()
        yield evmos_ws
    else:
        raise NotImplementedError


@pytest.fixture(scope="module", params=["silc", "silc-ws", "silc-rocksdb", "geth"])
def cluster(request, silc, evmos_rocksdb, geth):
    """
    run on silc, silc websocket,
    silc built with rocksdb (memIAVL + versionDB)
    and geth
    """
    provider = request.param
    if provider == "silc":
        yield silc
    elif provider == "silc-ws":
        evmos_ws = silc.copy()
        evmos_ws.use_websocket()
        yield evmos_ws
    elif provider == "geth":
        yield geth
    elif provider == "silc-rocksdb":
        yield evmos_rocksdb
    else:
        raise NotImplementedError


@pytest.fixture(scope="module", params=["silc", "silc-rocksdb"])
def evmos_cluster(request, silc, evmos_rocksdb):
    """
    run on silc default build &
    silc with rocksdb build and memIAVL + versionDB
    """
    provider = request.param
    if provider == "silc":
        yield silc
    elif provider == "silc-rocksdb":
        yield evmos_rocksdb
    else:
        raise NotImplementedError
