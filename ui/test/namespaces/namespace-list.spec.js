import {Provider} from "react-redux";
import NamespaceListContainer from '../../src/namespaces/namespace-list';
import {afterEach, beforeEach, describe, expect, it} from '@jest/globals';
import {cleanup, render} from '@testing-library/react';
import Chance from 'chance';
import {findAllNamespaces} from "../../src/namespaces/namespace-repository";

jest.mock('../../src/namespaces/namespace-repository');

function mockStoreWithState(state) {
    return {
        getState: () => state,
        dispatch: () => {
        },
        subscribe: () => {
        }
    };
}

function givenANameSpaceList(namespace) {
    return {
        namespaces: [{
            name: namespace
        }]
    };
}

describe('NamespaceList', () => {
    let chance;

    beforeEach(() => {
        chance = new Chance();
    });

    afterEach(() => {
        cleanup();
        jest.resetAllMocks();
    });

    it('should show a list of namespaces from the cluster', () => {
        const firstNamespace = chance.word();
        const expectedNamespaces = givenANameSpaceList(firstNamespace);
        const store = mockStoreWithState({
            app: expectedNamespaces
        });
        const namespaceListEl = render(<Provider store={store}><NamespaceListContainer/></Provider>);
        namespaceListEl.container.querySelectorAll('div > div').forEach((namespaceEl) => {
            expect(namespaceEl.textContent).toContain(firstNamespace);
        });
    });

    it('should use the namespace repository to find all namespaces', async () => {
        const store = mockStoreWithState({
            app: {
                namespaces: []
            }
        });
        findAllNamespaces.mockReturnValueOnce([1, 2, 3]);
        expect(findAllNamespaces).not.toHaveBeenCalled();

        render(<Provider store={store}><NamespaceListContainer/></Provider>);

        expect(findAllNamespaces).toHaveBeenCalled();
    });
});
