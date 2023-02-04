import {Provider} from "react-redux";
import NamespaceListContainer from '../../src/namespaces/namespace-list';
import {beforeEach, describe, it, expect} from '@jest/globals';
import {render} from '@testing-library/react';
import Chance from 'chance';

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

describe('NamespaceList', () => {
    let chance;

    beforeEach(() => {
        chance = new Chance();
    });

    it('should show a list of namespaces from the cluster', () => {
        const firstNamespace = chance.word();
        const expectedNamespaces = [{
            name: firstNamespace
        }];
        const store = mockStoreWithState({
            app: {
                namespaces: expectedNamespaces
            }
        })
        const namespaceListEl = render(<Provider store={store}><NamespaceListContainer/></Provider>);
        namespaceListEl.container.querySelectorAll('div > div').forEach((namespaceEl) => {
            expect(namespaceEl.textContent).toContain(firstNamespace);
        });
    });
});
