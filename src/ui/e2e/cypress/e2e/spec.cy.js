describe('item list', () => {
  it('shows item list header', () => {
    // arrange
    cy.visit('http://localhost:3001');

    // assert
    cy.get('h2').contains('Items List');
  })

  it('has 10 items', () => {
    // arrange
    cy.visit('http://localhost:3001');

    // assert
    cy.get('#item-list').children().should('have.length', 10)
  })
})

describe('total', ()=> {
  it('shows total header', () => {
    // arrange
    cy.visit('http://localhost:3001');

    // assert
    cy.get('h2').contains('Total');
  })

  it('has initial total of zero', () => {
    // arrange
    cy.visit('http://localhost:3001');

    // assert
    cy.get('#total-list').children().should('have.length', 0)
  })

  it('total list contains 1 element when item clicked', () => {
    // arrange
    cy.visit('http://localhost:3001');
    const item = cy.get('#item-list').children().first();

    // act
    item.click()

    // assert
    cy.get('#total-list').children().should('have.length', 1)
  })

  it('total list contains 2 elements when two items clicked', () => {
    // arrange
    cy.visit('http://localhost:3001');
    const firstItem = cy.get('#item-list').children().first();
    firstItem.click()
    const secondItem = firstItem.next();
    
    // act
    secondItem.click()
    
    // assert
    cy.get('#total-list').children().should('have.length', 2)
  })
})

describe('subtotal', ()=> {
  it('shows subtotal header', () => {
    // arrange
    cy.visit('http://localhost:3001');

    // assert
    cy.get('h3').contains('Subtotal');
  })

  it('has initial subtotal of zero', () => {
    // arrange
    cy.visit('http://localhost:3001');

    // assert
    cy.get('#subtotal').should('include.text', '$0.00')
  })

  it('subtotal $10.00 when purchasing 1 item', () => {
    // arrange
    cy.visit('http://localhost:3001');
    const item = cy.get('#item-list').children().first();

    // act
    item.click()

    // assert
    cy.get('#subtotal').should('include.text', '$10.00')
  })

  it('subtotal $20.00 when purchasing 2 items', () => {
    // arrange
    cy.visit('http://localhost:3001');
    const firstItem = cy.get('#item-list').children().first();
    firstItem.click()
    const secondItem = firstItem.next()

    // act
    secondItem.click()

    // assert
    cy.get('#subtotal').should('include.text', '$20.00')
  })
})