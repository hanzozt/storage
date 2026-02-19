/*
	Copyright NetFoundry, Inc.

	Licensed under the Apache License, Version 2.0 (the "License");
	you may not use this file except in compliance with the License.
	You may obtain a copy of the License at

	https://www.apache.org/licenses/LICENSE-2.0

	Unless required by applicable law or agreed to in writing, software
	distributed under the License is distributed on an "AS IS" BASIS,
	WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
	See the License for the specific language governing permissions and
	limitations under the License.
*/

package ast

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	ztql "github.com/hanzozt/storage/ztql"
	"runtime"
	"strings"
)

type LoggingListener struct {
	PrintRuleLocation bool
	PrintChildren     bool
}

var _ ztql.ZitiQlListener = (*LoggingListener)(nil)

func (l *LoggingListener) printRuleLocationWithSkip(s int) {
	if l.PrintRuleLocation {
		pc, _, _, _ := runtime.Caller(s)
		f := runtime.FuncForPC(pc)
		s := strings.Split(f.Name(), ".")
		println(s[len(s)-1])
	}
}

func (l *LoggingListener) printChildren(tree antlr.ParseTree) {
	if l.PrintChildren {
		fmt.Printf("children for: %s\n", tree.GetText())

		for i, c := range tree.GetChildren() {
			fmt.Printf("-- %d: %s\n", i, c.(antlr.ParseTree).GetText())
		}
	}
}

func (l *LoggingListener) printDebug(tree antlr.ParseTree) {
	l.printRuleLocationWithSkip(2)
	l.printChildren(tree)
}

func (l *LoggingListener) VisitTerminal(node antlr.TerminalNode) {
	l.printDebug(node)
}

func (l *LoggingListener) VisitErrorNode(node antlr.ErrorNode) {
	l.printDebug(node)
}

func (l *LoggingListener) EnterEveryRule(ctx antlr.ParserRuleContext) {
	l.printDebug(ctx)
}

func (l *LoggingListener) ExitEveryRule(ctx antlr.ParserRuleContext) {
	l.printDebug(ctx)
}

func (l *LoggingListener) EnterQueryStmt(c *ztql.QueryStmtContext) {
	l.printDebug(c)
}

func (l *LoggingListener) EnterSortByExpr(c *ztql.SortByExprContext) {
	l.printDebug(c)
}

func (l *LoggingListener) EnterSortFieldExpr(c *ztql.SortFieldExprContext) {
	l.printDebug(c)
}

func (l *LoggingListener) ExitQueryStmt(c *ztql.QueryStmtContext) {
	l.printDebug(c)
}

func (l *LoggingListener) ExitSortByExpr(c *ztql.SortByExprContext) {
	l.printDebug(c)
}

func (l *LoggingListener) ExitSortFieldExpr(c *ztql.SortFieldExprContext) {
	l.printDebug(c)
}

func (l *LoggingListener) EnterSkipExpr(c *ztql.SkipExprContext) {
	l.printDebug(c)
}

func (l *LoggingListener) EnterLimitExpr(c *ztql.LimitExprContext) {
	l.printDebug(c)
}

func (l *LoggingListener) ExitSkipExpr(c *ztql.SkipExprContext) {
	l.printDebug(c)
}

func (l *LoggingListener) ExitLimitExpr(c *ztql.LimitExprContext) {
	l.printDebug(c)
}

func (l *LoggingListener) EnterSetFunctionExpr(c *ztql.SetFunctionExprContext) {
	l.printDebug(c)
}

func (l *LoggingListener) EnterBinaryLhs(c *ztql.BinaryLhsContext) {
	l.printDebug(c)
}

func (l *LoggingListener) EnterStringArray(c *ztql.StringArrayContext) {
	l.printDebug(c)
}

func (l *LoggingListener) EnterNumberArray(c *ztql.NumberArrayContext) {
	l.printDebug(c)
}

func (l *LoggingListener) EnterDatetimeArray(c *ztql.DatetimeArrayContext) {
	l.printDebug(c)
}

func (l *LoggingListener) EnterEnd(c *ztql.EndContext) {
	l.printDebug(c)
}

func (l *LoggingListener) EnterGroup(c *ztql.GroupContext) {
	l.printDebug(c)
}

func (l *LoggingListener) EnterOrExpr(c *ztql.OrExprContext) {
	l.printDebug(c)
}

func (l *LoggingListener) EnterOperationOp(c *ztql.OperationOpContext) {
	l.printDebug(c)
}

func (l *LoggingListener) EnterAndExpr(c *ztql.AndExprContext) {
	l.printDebug(c)
}

func (l *LoggingListener) EnterInStringArrayOp(c *ztql.InStringArrayOpContext) {
	l.printDebug(c)
}

func (l *LoggingListener) EnterInNumberArrayOp(c *ztql.InNumberArrayOpContext) {
	l.printDebug(c)
}

func (l *LoggingListener) EnterInDatetimeArrayOp(c *ztql.InDatetimeArrayOpContext) {
	l.printDebug(c)
}

func (l *LoggingListener) EnterBetweenNumberOp(c *ztql.BetweenNumberOpContext) {
	l.printDebug(c)
}

func (l *LoggingListener) EnterBetweenDateOp(c *ztql.BetweenDateOpContext) {
	l.printDebug(c)
}

func (l *LoggingListener) EnterBinaryLessThanStringOp(c *ztql.BinaryLessThanStringOpContext) {
	l.printDebug(c)
}

func (l *LoggingListener) EnterBinaryGreaterThanStringOp(c *ztql.BinaryGreaterThanStringOpContext) {
	l.printDebug(c)
}

func (l *LoggingListener) ExitBinaryLessThanStringOp(c *ztql.BinaryLessThanStringOpContext) {
	l.printDebug(c)
}

func (l *LoggingListener) ExitBinaryGreaterThanStringOp(c *ztql.BinaryGreaterThanStringOpContext) {
	l.printDebug(c)
}

func (l *LoggingListener) EnterBinaryLessThanNumberOp(c *ztql.BinaryLessThanNumberOpContext) {
	l.printDebug(c)
}

func (l *LoggingListener) EnterBinaryLessThanDatetimeOp(c *ztql.BinaryLessThanDatetimeOpContext) {
	l.printDebug(c)
}

func (l *LoggingListener) EnterBinaryGreaterThanNumberOp(c *ztql.BinaryGreaterThanNumberOpContext) {
	l.printDebug(c)
}

func (l *LoggingListener) EnterBinaryGreaterThanDatetimeOp(c *ztql.BinaryGreaterThanDatetimeOpContext) {
	l.printDebug(c)
}

func (l *LoggingListener) EnterBinaryEqualToStringOp(c *ztql.BinaryEqualToStringOpContext) {
	l.printDebug(c)
}

func (l *LoggingListener) EnterBinaryEqualToNumberOp(c *ztql.BinaryEqualToNumberOpContext) {
	l.printDebug(c)
}

func (l *LoggingListener) EnterBinaryEqualToDatetimeOp(c *ztql.BinaryEqualToDatetimeOpContext) {
	l.printDebug(c)
}

func (l *LoggingListener) EnterBinaryEqualToBoolOp(c *ztql.BinaryEqualToBoolOpContext) {
	l.printDebug(c)
}

func (l *LoggingListener) ExitBinaryEqualToBoolOp(c *ztql.BinaryEqualToBoolOpContext) {
	l.printDebug(c)
}

func (l *LoggingListener) EnterBinaryEqualToNullOp(c *ztql.BinaryEqualToNullOpContext) {
	l.printDebug(c)
}

func (l *LoggingListener) EnterBinaryContainsOp(c *ztql.BinaryContainsOpContext) {
	l.printDebug(c)
}

func (l *LoggingListener) ExitSetFunctionExpr(c *ztql.SetFunctionExprContext) {
	l.printDebug(c)
}

func (l *LoggingListener) ExitBinaryLhs(c *ztql.BinaryLhsContext) {
	l.printDebug(c)
}

func (l *LoggingListener) ExitStringArray(c *ztql.StringArrayContext) {
	l.printDebug(c)
}

func (l *LoggingListener) ExitNumberArray(c *ztql.NumberArrayContext) {
	l.printDebug(c)
}

func (l *LoggingListener) ExitDatetimeArray(c *ztql.DatetimeArrayContext) {
	l.printDebug(c)
}

func (l *LoggingListener) ExitEnd(c *ztql.EndContext) {
	l.printDebug(c)
}

func (l *LoggingListener) ExitGroup(c *ztql.GroupContext) {
	l.printDebug(c)
}

func (l *LoggingListener) ExitOrExpr(c *ztql.OrExprContext) {
	l.printDebug(c)
}

func (l *LoggingListener) ExitOperationOp(c *ztql.OperationOpContext) {
	l.printDebug(c)
}

func (l *LoggingListener) ExitAndExpr(c *ztql.AndExprContext) {
	l.printDebug(c)
}

func (l *LoggingListener) ExitInStringArrayOp(c *ztql.InStringArrayOpContext) {
	l.printDebug(c)
}

func (l *LoggingListener) ExitInNumberArrayOp(c *ztql.InNumberArrayOpContext) {
	l.printDebug(c)
}

func (l *LoggingListener) ExitInDatetimeArrayOp(c *ztql.InDatetimeArrayOpContext) {
	l.printDebug(c)
}

func (l *LoggingListener) ExitBetweenNumberOp(c *ztql.BetweenNumberOpContext) {
	l.printDebug(c)
}

func (l *LoggingListener) ExitBetweenDateOp(c *ztql.BetweenDateOpContext) {
	l.printDebug(c)

}

func (l *LoggingListener) ExitBinaryLessThanNumberOp(c *ztql.BinaryLessThanNumberOpContext) {
	l.printDebug(c)
}

func (l *LoggingListener) ExitBinaryLessThanDatetimeOp(c *ztql.BinaryLessThanDatetimeOpContext) {
	l.printDebug(c)
}

func (l *LoggingListener) ExitBinaryGreaterThanNumberOp(c *ztql.BinaryGreaterThanNumberOpContext) {
	l.printDebug(c)
}

func (l *LoggingListener) ExitBinaryGreaterThanDatetimeOp(c *ztql.BinaryGreaterThanDatetimeOpContext) {
	l.printDebug(c)
}

func (l *LoggingListener) ExitBinaryEqualToStringOp(c *ztql.BinaryEqualToStringOpContext) {
	l.printDebug(c)
}

func (l *LoggingListener) ExitBinaryEqualToNumberOp(c *ztql.BinaryEqualToNumberOpContext) {
	l.printDebug(c)
}

func (l *LoggingListener) ExitBinaryEqualToDatetimeOp(c *ztql.BinaryEqualToDatetimeOpContext) {
	l.printDebug(c)
}

func (l *LoggingListener) ExitBinaryEqualToNullOp(c *ztql.BinaryEqualToNullOpContext) {
	l.printDebug(c)
}

func (l *LoggingListener) ExitBinaryContainsOp(c *ztql.BinaryContainsOpContext) {
	l.printDebug(c)
}

func (l *LoggingListener) EnterBoolConst(c *ztql.BoolConstContext) {
	l.printDebug(c)
}

func (l *LoggingListener) ExitBoolConst(c *ztql.BoolConstContext) {
	l.printDebug(c)
}

func (l *LoggingListener) EnterIsEmptyFunction(c *ztql.IsEmptyFunctionContext) {
	l.printDebug(c)
}

func (l *LoggingListener) EnterBoolSymbol(c *ztql.BoolSymbolContext) {
	l.printDebug(c)
}

func (l *LoggingListener) ExitIsEmptyFunction(c *ztql.IsEmptyFunctionContext) {
	l.printDebug(c)
}

func (l *LoggingListener) ExitBoolSymbol(c *ztql.BoolSymbolContext) {
	l.printDebug(c)
}

func (l *LoggingListener) EnterNotExpr(c *ztql.NotExprContext) {
	l.printDebug(c)
}

func (l *LoggingListener) ExitNotExpr(c *ztql.NotExprContext) {
	l.printDebug(c)
}

func (l *LoggingListener) EnterSetExpr(c *ztql.SetExprContext) {
	l.printDebug(c)
}

func (l *LoggingListener) EnterSubQuery(c *ztql.SubQueryContext) {
	l.printDebug(c)
}

func (l *LoggingListener) ExitSetExpr(c *ztql.SetExprContext) {
	l.printDebug(c)
}

func (l *LoggingListener) ExitSubQuery(c *ztql.SubQueryContext) {
	l.printDebug(c)
}
