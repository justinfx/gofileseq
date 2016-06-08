#include "common.h"
#include "../private/ranges.h"

#include <gtest/gtest.h>

#include <map>
#include <string>


// Range
//

struct TRangeInit {
    long start, end;
    int step;
    long rEnd;
    size_t rLen;
    std::string rStr;
};

TEST( TestRange, Init ) {
    typedef TRangeInit Case;
    std::vector<Case> cases;

    {Case t = {1, 10, 1, 10, 10, "1-10"}; cases.push_back(t);}
    {Case t = {10, 1, -1, 1, 10, "10-1"}; cases.push_back(t);}
    {Case t = {-10, 20, 1, 20, 31, "-10-20"}; cases.push_back(t);}
    {Case t = {20, -10, -1, -10, 31, "20--10"}; cases.push_back(t);}
    {Case t = {1, 10, 3, 10, 4, "1-10x3"}; cases.push_back(t);}
    {Case t = {10, -10, -3, -8, 7, "10--8x-3"}; cases.push_back(t);}
    {Case t = {-100, -33, 7, -37, 10, "-100--37x7"}; cases.push_back(t);}
    {Case t = {-33, -100, -7, -96, 10, "-33--96x-7"}; cases.push_back(t);}
    {Case t = {1, 10, 0, 10, 10, "1-10"}; cases.push_back(t);}
    {Case t = {10, 1, 0, 1, 10, "10-1"}; cases.push_back(t);}

    int step = 0;

    for (size_t i=0; i < cases.size(); ++i) {
        Case tt = cases[i];

        fileseq::Range r(tt.start, tt.end, tt.step);

        step = tt.step;
        if (tt.step == 0) {
            if (r.end() >= r.start()) {
                step = 1;
            } else {
                step = -1;
            }
        }
        EXPECT_EQ(step, r.step()) << "Step does not match (case: " << tt.rStr << ")";
        EXPECT_EQ(tt.start, r.start()) << "start does not match (case: " << tt.rStr << ")";
        EXPECT_EQ(tt.rEnd, r.end()) << "end does not match (case: " << tt.rStr << ")";
        EXPECT_EQ(tt.rLen, r.length()) << "Length does not match (case: " << tt.rStr << ")";
        EXPECT_STREQ(tt.rStr.c_str(), r.string().c_str());
    }
}


struct TRangeLookups {
    long start, end;
    int step;
    std::map<ssize_t, long> indexToValue;
    std::map<long, ssize_t> valueToIndex;

};

TEST( TestRange, Lookups ) {
    typedef TRangeLookups Case;
    std::vector<Case> cases;

    {
        Case t;
        t.start = 1;
        t.end = 10;
        t.step = 1;

        t.indexToValue[0] = 1;
        t.indexToValue[5] = 6;
        t.indexToValue[9] = 10;

        t.valueToIndex[1] = 0;
        t.valueToIndex[6] = 5;
        t.valueToIndex[10] = 9;
        t.valueToIndex[-1] = -1;
        t.valueToIndex[11] = -1;

        cases.push_back(t);
    }
    {
        Case t;
        t.start = 10;
        t.end = -10;
        t.step = -3;

        t.indexToValue[0] = 10;
        t.indexToValue[3] = 1;
        t.indexToValue[6] = -8;

        t.valueToIndex[10] = 0;
        t.valueToIndex[1] = 3;
        t.valueToIndex[-8] = 6;
        t.valueToIndex[3] = -1;
        t.valueToIndex[-7] = -1;

        cases.push_back(t);
    }

    typedef std::map<ssize_t, long> IV;
    typedef std::map<long, ssize_t> VI;

    fileseq::Status ok;
    long value;
    ssize_t idx;

    for (size_t i=0; i < cases.size(); ++i) {
        Case tt = cases[i];

        fileseq::Range r(tt.start, tt.end, tt.step);

        for (IV::iterator it = tt.indexToValue.begin(); it != tt.indexToValue.end(); ++it) {

            value = r.value(it->first, &ok);
            if (!ok) {
                ADD_FAILURE() << "Failed value lookup of idx " << it->first << ": " << ok;
                continue;
            }

            EXPECT_EQ(it->second, value) << "Given idx " << it->first;
        }

        for (VI::iterator it = tt.valueToIndex.begin(); it != tt.valueToIndex.end(); ++it) {
            idx = r.index(it->first);
            EXPECT_EQ(it->second, idx) << "Index not found for value " << it->first;
        }
    }
}

struct TRangeIterValues {
    long start, end;
    int step;
    std::vector<long> expected;
};

TEST( TestRange, IterValues ) {
    typedef TRangeIterValues Case;
    std::vector<Case> cases;

    {
        long arr[] = {-5, -4, -3, -2, -1, 0, 1, 2, 3, 4, 5};
        Case t = {-5, 5, 1, LONG_ARR_TO_VEC(arr)};
        cases.push_back(t);
    }
    {
        long arr[] = {-10, -7, -4, -1, 2, 5, 8};
        Case t = {-10, 10, 3, LONG_ARR_TO_VEC(arr)};
        cases.push_back(t);
    }
    {
        long arr[] = {20, 13, 6, -1, -8, -15, -22, -29};
        Case t = {20, -30, -7, LONG_ARR_TO_VEC(arr)};
        cases.push_back(t);
    }


    for (size_t i=0; i < cases.size(); ++i) {
        Case tt = cases[i];

        fileseq::Range r(tt.start, tt.end, tt.step);

        std::vector<long> actual;

        fileseq::RangeIterator rangeIt = r.iterValues();
        while (rangeIt.next()) {
            actual.push_back(*rangeIt);
        }
        EXPECT_EQ(tt.expected, actual);
    }
}

// Ranges
//

struct TRangesInit {
    std::vector< std::vector<long> > params;
    long rEnd;
    size_t rLen;
    std::string rStr;

    void addParams(long start, long end, long step) {
        std::vector<long> p;
        p.push_back(start);
        p.push_back(end);
        p.push_back(step);
        params.push_back(p);
    }
};

TEST( TestRanges, Init ) {
    typedef TRangesInit Case;
    std::vector<Case> cases;

    {
        Case t;
        t.rEnd = 0;
        t.rLen = 0;
        t.rStr = "";
        cases.push_back(t);
    }
    {
        Case t;
        t.rEnd = 28;
        t.rLen = 18;
        t.rStr = "1-10,13-14,17,16-28x3";
        t.addParams(1, 10, 1);
        t.addParams(13, 14, 1);
        t.addParams(17, 17, 1);
        t.addParams(1, 30, 3);
        cases.push_back(t);
    }
    {
        Case t;
        t.rEnd = 9;
        t.rLen = 21;
        t.rStr = "-5-5x2,10,-10--6,-4,-2,0,2,4,6-9";
        t.addParams(-5, 5, 2);
        t.addParams(10, 10, 1);
        t.addParams(-10, 10, 1);
        cases.push_back(t);
    }
    {
        Case t;
        t.rEnd = 19384765625;
        t.rLen = 19383275510;
        t.rStr = "1490116-19384765625";
        t.addParams(1490116, 19384765625, 1);
        cases.push_back(t);
    }

    for (size_t i=0; i < cases.size(); ++i) {
        Case tt = cases[i];

        fileseq::Ranges r;
        for (size_t i=0; i < tt.params.size(); ++i) {
            r.appendUnique(tt.params[i][0], tt.params[i][1], tt.params[i][2]);
        }

        long start = tt.params.empty() ? 0 : tt.params[0][0];

        EXPECT_EQ(start, r.start()) << "Start does not match";
        EXPECT_EQ(tt.rEnd, r.end()) << "End does not match";
        EXPECT_EQ(tt.rLen, r.length()) << "Length does not match";
        EXPECT_STREQ(tt.rStr.c_str(), r.string().c_str());
    }
}


struct TRangesMinMax {
    std::vector< std::vector<long> > params;
    long minVal, maxVal;

    void addParams(long start, long end, long step) {
        std::vector<long> p;
        p.push_back(start);
        p.push_back(end);
        p.push_back(step);
        params.push_back(p);
    }
};

TEST( TestRanges, MinMax ) {
    typedef TRangesMinMax Case;
    std::vector<Case> cases;

    {
        Case t;
        t.addParams(1, 10, 1);
        t.minVal = 1;
        t.maxVal = 10;
        cases.push_back(t);
    }
    {
        Case t;
        t.addParams(1, 10, 2);
        t.minVal = 1;
        t.maxVal = 9;
        cases.push_back(t);
    }
    {
        Case t;
        t.addParams(-6, 0, 2);
        t.addParams(2, 4, 1);
        t.addParams(10, 16, 3);
        t.minVal = -6;
        t.maxVal = 16;
        cases.push_back(t);
    }
    {
        Case t;
        t.addParams(16, 10, 3);
        t.addParams(4, 2, 1);
        t.addParams(0, -6, 2);
        t.minVal = -6;
        t.maxVal = 16;
        cases.push_back(t);
    }
    {
        Case t;
        t.addParams(4, 2, 1);
        t.addParams(0, -6, 2);
        t.addParams(16, 10, 3);
        t.addParams(100, -20, 5);
        t.addParams(2, 4, 1);
        t.minVal = -20;
        t.maxVal = 100;
        cases.push_back(t);
    }
    {
        Case t;
        t.addParams(1, 10, 2);
        t.addParams(50, 50, 1);
        t.addParams(60, 62, 2);
        t.addParams(70, 74, 2);
        t.minVal = 1;
        t.maxVal = 74;
        cases.push_back(t);
    }

    for (size_t i=0; i < cases.size(); ++i) {
        Case tt = cases[i];

        fileseq::Ranges r;
        for (size_t i=0; i < tt.params.size(); ++i) {
            r.appendUnique(tt.params[i][0], tt.params[i][1], tt.params[i][2]);
        }

        EXPECT_EQ(tt.minVal, r.min()) << "min val does not match for " << r.string();
        EXPECT_EQ(tt.maxVal, r.max()) << "max val does not match for " << r.string();
    }
}


struct TRangesLookups {
    std::vector< std::vector<long> > params;
    std::map<ssize_t, long> indexToValue;
    std::map<long, ssize_t> valueToIndex;

    void addParams(long start, long end, long step) {
        std::vector<long> p;
        p.push_back(start);
        p.push_back(end);
        p.push_back(step);
        params.push_back(p);
    }
};

TEST( TestRanges, Lookup ) {
    typedef TRangesLookups Case;
    std::vector<Case> cases;

    {
        Case t;
        t.addParams(1, 10, 1);
        t.addParams(13, 14, 1);
        t.addParams(17, 17, 1);
        t.addParams(1, 30, 3);

        t.indexToValue[0] = 1;
        t.indexToValue[10] = 13;
        t.indexToValue[12] = 17;
        t.indexToValue[17] = 28;

        t.valueToIndex[1] = 0;
        t.valueToIndex[13] = 10;
        t.valueToIndex[17] = 12;
        t.valueToIndex[28] = 17;
        t.valueToIndex[-1] = -1;
        t.valueToIndex[15] = -1;
        t.valueToIndex[30] = -1;

        cases.push_back(t);
    }
    {
        Case t;
        t.addParams(-5, 5, 2);
        t.addParams(10, 10, 1);
        t.addParams(-10, 10, 1);

        t.indexToValue[0] = -5;
        t.indexToValue[5] = 5;
        t.indexToValue[6] = 10;
        t.indexToValue[11] = -6;
        t.indexToValue[20] = 9;

        t.valueToIndex[-5] = 0;
        t.valueToIndex[5] = 5;
        t.valueToIndex[10] = 6;
        t.valueToIndex[-6] = 11;
        t.valueToIndex[9] = 20;
        t.valueToIndex[-11] = -1;
        t.valueToIndex[11] = -1;

        cases.push_back(t);
    }

    typedef std::map<ssize_t, long> IV;
    typedef std::map<long, ssize_t> VI;

    fileseq::Status ok;
    long value;
    ssize_t idx;

    for (size_t i=0; i < cases.size(); ++i) {
        Case tt = cases[i];

        fileseq::Ranges r;
        for (size_t i=0; i < tt.params.size(); ++i) {
            r.appendUnique(tt.params[i][0], tt.params[i][1], tt.params[i][2]);
        }

        for (IV::iterator it = tt.indexToValue.begin(); it != tt.indexToValue.end(); ++it) {

            value = r.value(it->first, &ok);
            if (!ok) {
                ADD_FAILURE() << "Failed value lookup of idx " << it->first << ": " << ok;
                continue;
            }

            EXPECT_EQ(it->second, value) << "Given idx " << it->first;
        }

        for (VI::iterator it = tt.valueToIndex.begin(); it != tt.valueToIndex.end(); ++it) {
            idx = r.index(it->first);
            EXPECT_EQ(it->second, idx) << "Index not found for value " << it->first;
        }
    }
}

struct TRangesIterValues {
    std::vector< std::vector<long> > params;
    std::vector<long> expected;

    void addParams(long start, long end, long step) {
        std::vector<long> p;
        p.push_back(start);
        p.push_back(end);
        p.push_back(step);
        params.push_back(p);
    }
};

TEST( TestRanges, IterValues ) {
    typedef TRangesIterValues Case;
    std::vector<Case> cases;

    {
        Case t;
        t.addParams(-5, 5, 2);
        t.addParams(10, 10, 1);
        t.addParams(-10, 10, 1);

        long arr[] = {
            -5, -3, -1, 1, 3, 5, 10, -10, -9, -8,
            -7, -6, -4, -2, 0, 2, 4, 6, 7, 8, 9,
        };
        t.expected = LONG_ARR_TO_VEC(arr);

        cases.push_back(t);
    }
    {
        Case t;
        t.addParams(-3, 3, 1);
        t.addParams(-5, 5, 1);

        long arr[] = {-3, -2, -1, 0, 1, 2, 3, -5, -4, 4, 5};
        t.expected = LONG_ARR_TO_VEC(arr);

        cases.push_back(t);
    }
    {
        Case t;
        t.addParams(10, 1, -2);
        t.addParams(-5, 10, 2);

        long arr[] = {10, 8, 6, 4, 2, -5, -3, -1, 1, 3, 5, 7, 9};
        t.expected = LONG_ARR_TO_VEC(arr);

        cases.push_back(t);
    }

    for (size_t i=0; i < cases.size(); ++i) {
        Case tt = cases[i];

        fileseq::Ranges r;
        for (size_t i=0; i < tt.params.size(); ++i) {
            r.appendUnique(tt.params[i][0], tt.params[i][1], tt.params[i][2]);
        }

        std::vector<long> actual;
        fileseq::RangesIterator rangeIt = r.iterValues();
        if (!rangeIt.isValid()) {
            ADD_FAILURE() << "RangesIterator reported invalid for: " << r.string();
            continue;
        }
        while (rangeIt.next()) {
            actual.push_back(*rangeIt);
        }
        EXPECT_EQ(tt.expected, actual);
    }
}


struct TRangesParamsString {
    std::vector< std::vector<long> > params;
    std::string expected;

    void addParams(long start, long end, long step) {
        std::vector<long> p;
        p.push_back(start);
        p.push_back(end);
        p.push_back(step);
        params.push_back(p);
    }
};

TEST( TestRanges, Inverted ) {
    typedef TRangesParamsString Case;
    std::vector<Case> cases;

    {
        Case t;
        t.addParams(1, 10, 1);
        t.expected = "";
        cases.push_back(t);
    }
    {
        Case t;
        t.addParams(1, 10, 2);
        t.expected = "2-8x2";
        cases.push_back(t);
    }
    {
        Case t;
        t.addParams(-6, 0, 2);
        t.addParams(2, 4, 1);
        t.addParams(10, 16, 3);
        t.expected = "-5-1x2,5-9,11-12,14-15";
        cases.push_back(t);
    }
    {
        Case t;
        t.addParams(16, 10, 3);
        t.addParams(4, 2, 1);
        t.addParams(0, -6, 2);
        t.expected = "-5-1x2,5-9,11-12,14-15";
        cases.push_back(t);
    }
    {
        Case t;
        t.addParams(4, 2, 1);
        t.addParams(0, -6, 2);
        t.addParams(16, 10, 3);
        t.addParams(2, 4, 1);
        t.expected = "-5-1x2,5-9,11-12,14-15";
        cases.push_back(t);
    }
    {
        Case t;
        t.addParams(1, 10, 2);
        t.addParams(50, 50, 1);
        t.addParams(60, 62, 2);
        t.addParams(70, 74, 2);
        t.expected = "2-10x2,11-49,51-59,61-63x2,64-69,71-73x2";
        cases.push_back(t);
    }

    for (size_t i=0; i < cases.size(); ++i) {
        Case tt = cases[i];

        fileseq::Ranges r;
        for (size_t i=0; i < tt.params.size(); ++i) {
            r.appendUnique(tt.params[i][0], tt.params[i][1], tt.params[i][2]);
        }

        std::cerr << "DEBUG: for seq: " << r.string()
                  << ", min = " << r.min() << ", max = " << r.max()
                  << std::endl;

        fileseq::Ranges inverted;
        r.inverted(inverted);

        EXPECT_STREQ(
            tt.expected.c_str(),
            inverted.string().c_str()) << "inverted frame range doesn't match";
    }
}

TEST( TestRanges, Normalized ) {
    typedef TRangesParamsString Case;
    std::vector<Case> cases;

    {
        Case t;
        t.addParams(1, 1, 1);
        t.addParams(2, 2, 1);
        t.addParams(3, 3, 1);
        t.addParams(4, 4, 1);
        t.addParams(5, 5, 1);
        t.expected = "1-5";
        cases.push_back(t);
    }
    {
        Case t;
        t.addParams(5, 5, 1);
        t.addParams(1, 1, 1);
        t.addParams(4, 4, 1);
        t.addParams(2, 2, 1);
        t.addParams(3, 3, 1);
        t.addParams(1, 1, 1);
        t.addParams(1, 1, 1);
        t.expected = "1-5";
        cases.push_back(t);
    }
    {
        Case t;
        t.addParams(1, 5, 1);
        t.addParams(6, 6, 1);
        t.addParams(7, 7, 1);
        t.addParams(8, 8, 1);
        t.addParams(9, 9, 1);
        t.addParams(10, 10, 1);
        t.addParams(11, 20, 1);
        t.expected = "1-20";
        cases.push_back(t);
    }

    for (size_t i=0; i < cases.size(); ++i) {
        Case tt = cases[i];

        fileseq::Ranges r;
        for (size_t i=0; i < tt.params.size(); ++i) {
            r.append(tt.params[i][0], tt.params[i][1], tt.params[i][2]);
        }

        fileseq::Ranges normalized;
        r.normalized(normalized);
        EXPECT_EQ(tt.expected, normalized.string()) << "normalized frame range doesn't match";
    }
}
